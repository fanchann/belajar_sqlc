package repositories

import "latihan_sqlc/internal/models/domain"

type IBoookRepositories interface {
	AddBook(book domain.Books) domain.Books
	GetBookById(id int) (domain.Books, error)
	UpdateBookById(book domain.Books) domain.Books
	GetAllBook() []domain.Books
	DeleteBookById(book domain.Books) error
}
