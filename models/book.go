package models

type Book struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	AuthorID   int       `json:"author_id"`
	CategoryID int       `json:"category_id"`
	Price      float64   `json:"price"`
	Author     *Author   `json:"author,omitempty"`   // Связь с автором
	Category   *Category `json:"category,omitempty"` // Связь с категорией
}
