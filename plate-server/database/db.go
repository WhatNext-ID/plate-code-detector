package database

import (
	"fmt"
	"log"
	"os"
	"server/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	databaseURL := os.Getenv("DATABASE_URL")
	fmt.Println("DATABASE_URL:", databaseURL)

	db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Terjadi Kesalahan saat koneksi ke db :", err)
	}

	fmt.Println("Berhasil tersambung ke db")
	db.AutoMigrate(models.MStatusKendaraan{}, models.MAdmin{})
}

func GetDB() *gorm.DB {
	return db
}
