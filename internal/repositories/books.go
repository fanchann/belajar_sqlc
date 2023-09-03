package repositories

import "latihan_sqlc/internal/models"

type IBoookRepositories interface {
	AddBook(book models.Books) models.Books
	GetBookById(id int) (models.Books, error)
	UpdateBookById(book models.Books) models.Books
	GetAllBook() []models.Books
	DeleteBookById(id int) error
}
