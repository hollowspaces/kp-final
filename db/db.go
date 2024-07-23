package db

import (
	"kp-final/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	// Konfigurasi koneksi ke database PostgreSQL
	dsn := "host=dpg-cqfp2kaju9rs73c1dr2g-a user=postgresql_puji_user password=88UMCaUBA2LdzGINsOqtDSFIjeXxzP1T dbname=postgresql_puji port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var err error

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
