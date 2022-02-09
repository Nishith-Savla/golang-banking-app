package dto

import (
	"github.com/Nishith-Savla/golang-banking-lib/errs"
	"strings"
)

const WITHDRAWAL = "withdrawal"
const DEPOSIT = "deposit"

type TransactionRequest struct {
	AccountId  string  `json:"account_id"`
	Amount     float64 `json:"amount"`
	Type       string  `json:"transaction_type"`
	Date       string  `json:"transaction_date"`
	CustomerId string  `json:"-"`
}

func (r TransactionRequest) IsWithdrawal() bool {
	return strings.ToLower(r.Type) == WITHDRAWAL
}

func (r TransactionRequest) IsDeposit() bool {
	return strings.ToLower(r.Type) == DEPOSIT
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < 1 {
		return errs.NewValidationError("To perform a new transaction, amount should be greater than ₹0.")
	}
	if !r.IsWithdrawal() && !r.IsDeposit() {
		return errs.NewValidationError("Transaction type should be withdrawal or deposit")
	}

	return nil
}

type TransactionResponse struct {
	Id        string  `json:"transaction_id"`
	AccountId string  `json:"account_id"`
	Amount    float64 `json:"new_balance"`
	Type      string  `json:"transaction_type"`
	Date      string  `json:"transaction_date"`
}
