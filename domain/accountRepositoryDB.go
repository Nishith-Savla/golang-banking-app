package domain

import (
	"database/sql"
	"github.com/Nishith-Savla/golang-banking-lib/errs"
	"github.com/Nishith-Savla/golang-banking-lib/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	insertSql := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"
	result, err := d.client.Exec(insertSql, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last inert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func (d AccountRepositoryDb) FindBy(id string) (*Account, *errs.AppError) {
	selectSql := "SELECT * FROM accounts WHERE account_id = ?"

	var a Account
	err := d.client.Get(&a, selectSql, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("account not found")
		}
		logger.Error("Error while scanning account " + id + ": " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	insertSql := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)"
	result, err := tx.Exec(insertSql, t.AccountId, t.Amount, t.Type, t.Date)
	if err != nil {
		logger.Error("Error while inserting a new transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while retrieving the last insert id for transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	updateSql := "UPDATE accounts SET amount = amount + ? WHERE account_id = ?"
	if t.IsWithdrawal() {
		updateSql = strings.Replace(updateSql, "+", "-", 1)
	}
	_, err = tx.Exec(updateSql, t.Amount, t.AccountId)

	if err != nil {
		_ = tx.Rollback()
		logger.Error("Error while updating the amount after saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
		logger.Error("Error while committing transaction for bank account" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	account, appError := d.FindBy(t.AccountId)
	if appError != nil {
		return nil, appError
	}
	t.Id = strconv.FormatInt(transactionId, 10)

	// updating transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{dbClient}
}
