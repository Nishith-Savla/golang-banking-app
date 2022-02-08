package app

import (
	"encoding/json"
	"github.com/Nishith-Savla/golang-banking-app/logger"
	"github.com/Nishith-Savla/golang-banking-app/service"
	"github.com/gorilla/mux"
	"net/http"
)

//type Customer struct {
//	Name    string `json:"full_name" xml:"full_name"`
//	City    string `json:"city" xml:"city"`
//	Zipcode string `json:"zip_code" xml:"zip_code"`
//}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, err.AsMessage())
		return
	}

	writeJSONResponse(w, http.StatusOK, customers)

	// OLD: allow xml or json response
	//contentType := r.Header.Get("Content-Type")
	//
	//if contentType == "application/xml" {
	//	w.Header().Add("Content-Type", contentType)
	//	xml.NewEncoder(w).Encode(customers)
	//} else {
	//	w.Header().Add("Content-Type", "application/json")
	//	json.NewEncoder(w).Encode(customers)
	//}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeJSONResponse(w, err.Code, err.AsMessage())
		return
	}
	writeJSONResponse(w, http.StatusOK, &customer)
}

func writeJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Panic(err.Error())
	}
}
