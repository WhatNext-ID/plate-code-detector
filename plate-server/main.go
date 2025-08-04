package main

import (
	"fmt"
	"log"
	"os"
	"plate-server/database"
	"plate-server/router"

	"github.com/joho/godotenv"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		database.StartDB()
		database.RunMigrations()
		fmt.Println("Migrations completed successfully.")
		return
	}

	database.StartDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var PORT = os.Getenv("PORT")

	r := router.StartApp()
	r.Run(":" + PORT)
}
