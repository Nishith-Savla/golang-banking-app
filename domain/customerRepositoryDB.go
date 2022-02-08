package domain

import (
	"database/sql"
	"github.com/Nishith-Savla/golang-banking-app/errs"
	"github.com/Nishith-Savla/golang-banking-app/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)

	findAllSql := "SELECT * FROM customers"

	if status == "" {
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql += " WHERE status=?"
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// Not needed after changing client to sqlx.DB and using Select instead of Query
	//err = sqlx.StructScan(rows, &customers)
	//if err != nil {
	//	logger.Error("Error while scanning customers: " + err.Error())
	//	return nil, errs.NewUnexpectedError("unexpected database error")
	//}

	// Not needed after adding sqlx.StructScan
	//for rows.Next() {
	//	var c Customer
	//	err = rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	//	if err != nil {
	//		logger.Error("Error while scanning customers: " + err.Error())
	//		return nil, errs.NewUnexpectedError("unexpected database error")
	//	}
	//	customers = append(customers, c)
	//}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT * FROM customers WHERE customer_id=?"

	var c Customer

	//row := d.client.QueryRow(customerSql, id)
	//err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while scanning customer " + id + ": " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}
