package service

import (
	"github.com/go-playground/validator/v10"

	"latihan_sqlc/exception"
	"latihan_sqlc/internal/models/domain"
	"latihan_sqlc/internal/models/web"
	"latihan_sqlc/internal/repositories"
	"latihan_sqlc/utils"
)

var (
	errValidate error
)

type BookServiceImpl struct {
	Repo     *repositories.BookRepositoriesImpl
	Validate *validator.Validate
}

func NewBookServiceImpl(Repo *repositories.BookRepositoriesImpl, Validate *validator.Validate) *BookServiceImpl {
	return &BookServiceImpl{Repo: Repo, Validate: Validate}
}

func (s *BookServiceImpl) AddNewBook(request web.BookCreateForm) web.BookResponse {
	errValidate = s.Validate.Struct(&request)
	if errValidate != nil {
		panic(exception.NewValidationError(errValidate))
	}

	book := domain.Books{Title: request.Title, Author: request.Author}

	book = s.Repo.AddBook(book)

	return utils.ToBookResponse(book)
}

func (s *BookServiceImpl) UpdateBook(request web.BookUpdateForm) web.BookResponse {
	errValidate = s.Validate.Struct(&request)
	if errValidate != nil {
		panic(exception.NewValidationError(errValidate))
	}

	_, errNotFound := s.Repo.GetBookById(request.Id)
	if errNotFound != nil {
		panic(exception.NewNotFoundErr(errNotFound))
	}

	updatedBook := s.Repo.UpdateBookById(domain.Books(request))
	return utils.ToBookResponse(updatedBook)
}

func (s *BookServiceImpl) Delete(bookId int) {
	bookData, errNotFound := s.Repo.GetBookById(bookId)
	if errNotFound != nil {
		panic(exception.NewNotFoundErr(errNotFound))
	}

	if errWhileDete := s.Repo.DeleteBookById(bookData); errWhileDete != nil {
		panic(exception.NewNotFoundErr(errWhileDete))
	}

}

func (s *BookServiceImpl) FindBookById(bookId int) web.BookResponse {
	bookData, errNotFound := s.Repo.GetBookById(bookId)
	if errNotFound != nil {
		panic(exception.NewNotFoundErr(errNotFound))
	}

	return utils.ToBookResponse(bookData)
}

func (s *BookServiceImpl) GetAllBooks() []web.BookResponse {
	var booksResponse []web.BookResponse

	booksData := s.Repo.GetAllBook()

	for _, book := range booksData {
		booksResponse = append(booksResponse, utils.ToBookResponse(book))
	}

	return booksResponse
}
