package models

// Struktur data Book digunakan untuk merepresentasikan informasi buku.
type Book struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	BookTitle     string `json:"book_title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
	Publisher     string `json:"publisher"`
	ISBN          string `json:"isbn"`
	Page          int    `json:"page"`
}
