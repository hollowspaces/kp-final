package controllers

import (
	database "kp-final/db"
	"kp-final/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooks(c echo.Context) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(http.StatusOK, books)
}

func GetBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Buku tidak ditemukan")
	}
	return c.JSON(http.StatusOK, book)
}

func CreateBook(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, "Permintaan tidak valid")
	}
	database.DB.Create(book)
	return c.JSON(http.StatusCreated, book)
}

func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Buku tidak ditemukan")
	}
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, "Permintaan tidak valid")
	}
	database.DB.Save(&book)
	return c.JSON(http.StatusOK, book)
}

func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Buku tidak ditemukan")
	}
	database.DB.Delete(&book)
	return c.NoContent(http.StatusNoContent)
}
