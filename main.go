package main

import (
	_ "myGram/docs"
	"myGram/handler"
)

// @title MyGram App
// @version 1.0
// @description Final Project 1

// @host localhost:8080
// @BasePath /API/v1
func main() {

	// @securityDefinitions.apikey BearerAuth
	// @in header
	// @name Authorization
	// @description Type "Bearer" followed by a space and token you got from the User Login api.
	handler.StartApp()
}
