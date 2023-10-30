package db

import (
	"kp-final/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root@tcp(127.0.0.1:3306)/kp_final_puji?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal terhubung ke database: " + err.Error())
	}
	err = DB.AutoMigrate(&models.Book{}) // Migrasi model Book ke database
	if err != nil {
		panic("Gagal melakukan migrasi database: " + err.Error())
	}
}
