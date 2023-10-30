package controllers

import (
	"kp-final/models"
	"net/http"

	database "kp-final/db"

	"github.com/labstack/echo/v4"
)

// Fungsi getBooks digunakan untuk mengambil semua data buku dari database
// dan mengirimkannya sebagai respon JSON.
func GetBooks(c echo.Context) error {
	var books []models.Book
	database.DB.Find(&books)
	return c.JSON(http.StatusOK, books)
}

// Fungsi getBook digunakan untuk mengambil data buku berdasarkan ID
// dari database dan mengirimkannya sebagai respon JSON.
func GetBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Buku tidak ditemukan")
	}
	return c.JSON(http.StatusOK, book)
}

// Fungsi createBook digunakan untuk membuat buku baru dengan data
// yang diberikan dalam permintaan dan menyimpannya ke dalam database.
func CreateBook(c echo.Context) error {
	book := new(models.Book)
	if err := c.Bind(book); err != nil {
		return c.JSON(http.StatusBadRequest, "Permintaan tidak valid")
	}
	database.DB.Create(book)
	return c.JSON(http.StatusCreated, book)
}

// Fungsi updateBook digunakan untuk memperbarui buku berdasarkan ID
func UpdateBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Buku tidak ditemukan")
	}

	// Bind data terupdate ke dalam model book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, "Permintaan tidak valid")
	}

	// Memperbaharui data buku ke dalam database
	if err := database.DB.Save(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal memperbarui buku")
	}

	return c.JSON(http.StatusOK, book)
}

// Fungsi deleteBook digunakan untuk menghapus buku berdasarkan ID dari database.
func DeleteBook(c echo.Context) error {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.First(&book, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, "Buku tidak ditemukan")
	}

	if err := database.DB.Delete(&book).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, "Gagal menghapus buku")
	}

	return c.JSON(http.StatusOK, "Buku berhasil dihapus")
}
