package main

import (
	"log"
	"usrmanagement/configs"
	"usrmanagement/models"
	"usrmanagement/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configs.ConnectDB()
	models.MigrateDB()

	r := routes.SetupRouter() // Use the router returned by SetupRouter

	errr := r.Run(":8081")
	if errr != nil {
		log.Fatal("Error starting the server:", errr)
	}
}
