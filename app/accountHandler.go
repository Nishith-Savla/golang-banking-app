package app

import (
	"encoding/json"
	"github.com/Nishith-Savla/golang-banking-app/dto"
	"github.com/Nishith-Savla/golang-banking-app/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (h AccountHandler) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerId = customerId
	account, appError := h.service.NewAccount(request)
	if appError != nil {
		writeJSONResponse(w, appError.Code, appError.AsMessage())
		return
	}
	writeJSONResponse(w, http.StatusCreated, account)
}

func (h AccountHandler) makeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]

	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeJSONResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.AccountId = accountId
	request.CustomerId = customerId

	transaction, appError := h.service.MakeTransaction(request)
	if appError != nil {
		writeJSONResponse(w, appError.Code, appError.AsMessage())
		return
	}

	writeJSONResponse(w, http.StatusOK, transaction)
}
