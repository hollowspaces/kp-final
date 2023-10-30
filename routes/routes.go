// project-directory/routes/routes.go
package routes

import (
	"kp-final/controllers"

	"github.com/labstack/echo/v4"
)

// SetupRoutes adalah fungsi yang digunakan untuk menambahkan rute-rute ke dalam server Echo.
func SetupRoutes(e *echo.Echo) {
	e.GET("/books", controllers.GetBooks)
	e.GET("/books/:id", controllers.GetBook)
	e.POST("/books", controllers.CreateBook)
	e.PUT("/books/:id", controllers.UpdateBook)
	e.DELETE("/books/:id", controllers.DeleteBook)
}
