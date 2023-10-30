// project-directory/main.go
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

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	routes.SetupRoutes(e)
	e.Start(":8080")
}
