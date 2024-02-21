package model

import (
	"github.com/jinzhu/gorm"
	"github.com/toliveiradev/learning-go/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title     string `gorm:""json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID = ?", id).Find(&book)
	return &book, db
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(book)
	return book
}
