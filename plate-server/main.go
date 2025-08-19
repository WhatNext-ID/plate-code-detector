package main

import (
	"fmt"
	"os"
	"plate-server/database"
	"plate-server/router"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		database.StartDB()
		database.RunMigrations()
		fmt.Println("Migrations completed successfully.")
		return
	}

	database.StartDB()

	var PORT = os.Getenv("PORT")

	r := router.StartApp()
	r.Run(":" + PORT)
}
