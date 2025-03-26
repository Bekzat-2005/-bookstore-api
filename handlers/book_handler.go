package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bookstore/models"
	"github.com/gorilla/mux"
)

var books = []models.Book{
	{ID: 1, Title: "Abai joly", AuthorID: 1, CategoryID: 1, Price: 29.99},
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	for i, book := range books {
		for _, author := range authors {
			if book.AuthorID == author.ID {
				books[i].Author = &author
			}
		}
		for _, category := range categories {
			if book.CategoryID == category.ID {
				books[i].Category = &category
			}
		}
	}

	json.NewEncoder(w).Encode(books)
}

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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	book.ID = len(books) + 1
	books = append(books, book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

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
