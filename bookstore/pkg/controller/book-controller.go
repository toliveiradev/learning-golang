package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/toliveiradev/learning-go/bookstore/pkg/model"
	"github.com/toliveiradev/learning-go/bookstore/pkg/utils"
)

var NewBook model.Book

const jsonContentType = "application/json"

const contentTypeHeader = "Content-Type"

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := model.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book, _ := model.GetBookById(ID)
	res, _ := json.Marshal(book)
	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &model.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := &model.Book{}
	utils.ParseBody(r, book)
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	bookData, db := model.GetBookById(ID)
	if book.Name != "" {
		bookData.Name = book.Name
	}
	if book.Author != "" {
		bookData.Author = book.Author
	}
	if book.Publisher != "" {
		bookData.Publisher = book.Publisher
	}
	db.Save(&bookData)
	res, _ := json.Marshal(bookData)
	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	book := model.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set(contentTypeHeader, jsonContentType)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
