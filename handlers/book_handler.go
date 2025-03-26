package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bookstore/models"
	"github.com/gorilla/mux"
)

var books = []models.Book{
	{ID: 1, Title: "Go Programming", AuthorID: 1, CategoryID: 1, Price: 29.99},
}

// GetBooks возвращает список всех книг
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetBook возвращает книгу по ID
func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, book := range books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// CreateBook добавляет новую книгу
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = len(books) + 1
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook обновляет данные книги
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, book := range books {
		if book.ID == id {
			json.NewDecoder(r.Body).Decode(&books[i])
			books[i].ID = id
			json.NewEncoder(w).Encode(books[i])
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// DeleteBook удаляет книгу
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
