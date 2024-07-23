package db

import (
	"kp-final/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"fmt"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDB() error {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Konfigurasi koneksi ke database PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	// Membuka koneksi ke database menggunakan driver PostgreSQL
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database: " + err.Error())
	}

	// Melakukan migrasi model ke database
	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		panic("Gagal melakukan migrasi database: " + err.Error())
	}
	return nil
}
