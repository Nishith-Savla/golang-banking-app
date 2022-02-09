package main

import (
	"github.com/Nishith-Savla/golang-banking-app/app"
	"github.com/Nishith-Savla/golang-banking-lib/logger"
)

func main() {
	logger.Info("Starting our application...")
	app.Start()

}
