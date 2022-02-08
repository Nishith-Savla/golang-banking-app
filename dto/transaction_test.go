package dto

import (
	"net/http"
	"testing"
)

func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawal(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		Type:   "invalid transaction type",
		Amount: 1,
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Transaction type should be withdrawal or deposit" {
		t.Error("Invalid message while testing transaction type")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid Code while testing transaction type")
	}
}

func Test_should_return_error_when_amount_is_less_than_zero(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		Amount: -100,
		Type:   DEPOSIT,
	}

	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "To perform a new transaction, amount should be greater than ₹0." {
		t.Error("Invalid message while testing transaction amount")
	}
	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid Code while testing transaction amount")
	}
}
