package handlers

import (
	"bookstore/models"
	"encoding/json"
	"net/http"
)

var categories = []models.Category{
	{ID: 1, Name: "Fiction"},
	{ID: 2, Name: "Science"},
}

// GetCategories возвращает список всех категорий
func GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categories)
}

// CreateCategory добавляет новую категорию
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)
	category.ID = len(categories) + 1
	categories = append(categories, category)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(category)
}
