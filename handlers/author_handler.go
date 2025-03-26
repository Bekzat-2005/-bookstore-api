package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

var authors = []models.Author{
	{ID: 1, Name: "John Doe"},
	{ID: 2, Name: "Jane Smith"},
}

// GetAuthors возвращает список всех авторов
func GetAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

// CreateAuthor добавляет нового автора
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)
	author.ID = len(authors) + 1
	authors = append(authors, author)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}
