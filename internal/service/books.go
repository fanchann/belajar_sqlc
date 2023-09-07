package service

import "latihan_sqlc/internal/models/web"

type IBoookService interface {
	AddNewBook(request web.BookCreateForm) web.BookResponse
	UpdateBook(request web.BookUpdateForm) web.BookResponse
	Delete(bookId int)
	FindBookById(bookId int) web.BookResponse
	GetAllBooks() []web.BookResponse
}
