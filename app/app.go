package app

import (
	"fmt"
	"github.com/Nishith-Savla/golang-banking-app/domain"
	"github.com/Nishith-Savla/golang-banking-app/service"
	"github.com/Nishith-Savla/golang-banking-lib/logger"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"os"
	"time"
)

func sanityCheck() {
	envProps := []string{
		"SERVER_ADDRESS",
		"SERVER_PORT",
		"AUTH_SERVER_ADDRESS",
		"AUTH_SERVER_PORT",
		"DB_USER",
		"DB_PASSWORD",
		"DB_ADDRESS",
		"DB_PORT",
		"DB_NAME",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			logger.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func Start() {
	sanityCheck()

	// router := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	dbClient := getDbClient()

	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	//transactionRepositoryDb := domain.NewTransactionRepositoryDb(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}

	// define routes
	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.
		HandleFunc("/customer/{customer_id:\\d+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customer/{customer_id:\\d+}/account", ah.newAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customer/{customer_id:\\d+}/account/{account_id:\\d+}", ah.makeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	am := AuthMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	//router.HandleFunc("/customer", createCustomer).Methods(http.MethodPost)
	//router.HandleFunc("/customer/{customer_id:\\d+}", getCustomer).Methods(http.MethodGet)

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	logger.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%v", address, port), router).Error())
}

func getDbClient() *sqlx.DB {

	// initiating connection with database
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddress := os.Getenv("DB_ADDRESS")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddress, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
