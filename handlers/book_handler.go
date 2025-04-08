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

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем параметры из query
	categoryIDParam := r.URL.Query().Get("category_id")
	authorIDParam := r.URL.Query().Get("author_id")
	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	// Парсим их в int, добавим базовые проверки
	categoryID, _ := strconv.Atoi(categoryIDParam)
	authorID, _ := strconv.Atoi(authorIDParam)
	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10 // по умолчанию 10, если не указали
	}

	// Фильтруем книги
	var filteredBooks []models.Book
	for _, book := range books {
		// Фильтр по category_id и author_id (если они указаны)
		if (categoryID == 0 || book.CategoryID == categoryID) &&
			(authorID == 0 || book.AuthorID == authorID) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	// Пагинация
	start := (page - 1) * limit
	end := start + limit

	if start > len(filteredBooks) {
		start = len(filteredBooks)
	}
	if end > len(filteredBooks) {
		end = len(filteredBooks)
	}

	// Результатко
	json.NewEncoder(w).Encode(filteredBooks[start:end])
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
