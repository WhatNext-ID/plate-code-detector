package database

import (
	"fmt"
	"log"
	"os"
	"server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		log.Fatal("Terjadi Kesalahan saat koneksi ke db :", err)
	}

	fmt.Println("Berhasil tersambung ke db")
	db.AutoMigrate(models.MStatusKendaraan{})
}

func GetDB() *gorm.DB {
	return db
}
