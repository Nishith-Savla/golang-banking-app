package domain

import (
	"github.com/Nishith-Savla/golang-banking-app/dto"
	"github.com/Nishith-Savla/golang-banking-app/errs"
)

type Transaction struct {
	Id        string  `db:"transaction_id"`
	AccountId string  `db:"account_id"`
	Amount    float64 `db:"amount"`
	Type      string  `db:"transaction_type"`
	Date      string  `db:"transaction_date"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.Type == "withdrawal"
}

func (t Transaction) ToNewTransactionResponseDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		Id:        t.Id,
		AccountId: t.AccountId,
		Amount:    t.Amount,
		Type:      t.Type,
		Date:      t.Date,
	}
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}
