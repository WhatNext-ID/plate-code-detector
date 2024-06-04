package main

import (
	"log"
	"os"
	"server/database"
	"server/router"

	"github.com/joho/godotenv"
)

func main() {
	database.StartDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var PORT = os.Getenv("PORT")

	r := router.StartApp()
	r.Run(":" + PORT)
}
