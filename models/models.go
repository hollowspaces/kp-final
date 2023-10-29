package model

import (
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	BookTitle     string `json:"book_title"`
	Author        string `json:"author"`
	PublishedYear int    `json:"published_year"`
	Publisher     string `json:"publisher"`
	ISBN          string `json:"isbn"`
	Page          int    `json:"page"`
}
