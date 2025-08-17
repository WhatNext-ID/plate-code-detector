package database

import (
	"fmt"
	"log"
	"os"
	platecode "plate-server/models/plate-code"
	"plate-server/models/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	var err error

	databaseURL := os.Getenv("DATABASE_URL")
	db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Terjadi Kesalahan saat koneksi ke db :", err)
	}

	fmt.Println("Berhasil tersambung ke db")
}

func RunMigrations() {
	if db == nil {
		StartDB()
	}

	err := db.AutoMigrate(
		&user.UserRole{},
		&user.User{},
		&platecode.RegionPlateCode{},
		&platecode.RegisterPlateCode{},
		&platecode.VehicleEngine{},
		&platecode.VehicleType{},
		&platecode.VehicleCategory{},
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	fmt.Println("Migrations completed successfully.")
}

func GetDB() *gorm.DB {
	return db
}
