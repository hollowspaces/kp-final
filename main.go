package main

import (
	"kp-final/db"
	"kp-final/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Inisialisasi database
	if err := db.InitDB(); err != nil {
		panic("Gagal menginisialisasi database: " + err.Error())
	}

	e := echo.New()             // Inisialisasi server Echo
	e.Use(middleware.Logger())  // Middleware
	e.Use(middleware.Recover()) // Inisialisasi rute
	routes.SetupRoutes(e)

	e.Start(":8080") // Start server pada port 8080
}
