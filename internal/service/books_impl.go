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
	repo     repositories.IBoookRepositories
	validate *validator.Validate
}

func (s *BookServiceImpl) AddNewBook(request web.BookCreateForm) web.BookResponse {
	errValidate = s.validate.Struct(request)
	utils.PanicIfError(errValidate)

	book := domain.Books{Title: request.Title, Author: request.Author}

	book = s.repo.AddBook(book)

	return utils.ToBookResponse(book)
}

func (s *BookServiceImpl) UpdateBook(request web.BookUpdateForm) web.BookResponse {
	errValidate = s.validate.Struct(request)
	utils.PanicIfError(errValidate)

	bookData, errNotFound := s.repo.GetBookById(request.Id)
	if errNotFound != nil {
		panic(exception.NewNotFoundErr(errNotFound.Error()))
	}

	bookData.Author = request.Author
	bookData.Title = request.Title

	updatedBook := s.repo.UpdateBookById(bookData)
	return utils.ToBookResponse(updatedBook)
}

func (s *BookServiceImpl) Delete(bookId int) {
	bookData, errNotFound := s.repo.GetBookById(bookId)
	if errNotFound != nil {
		panic(exception.NewNotFoundErr(errNotFound.Error()))
	}

	if errWhileDete := s.repo.DeleteBookById(bookData); errWhileDete != nil {
		panic(exception.NewNotFoundErr(errWhileDete.Error()))
	}

}

func (s *BookServiceImpl) FindBookById(bookId int) web.BookResponse {
	bookData, errNotFound := s.repo.GetBookById(bookId)
	if errNotFound != nil {
		panic(exception.NewNotFoundErr(errNotFound.Error()))
	}

	return utils.ToBookResponse(bookData)
}

func (s *BookServiceImpl) GetAllBooks() []web.BookResponse {
	var booksResponse []web.BookResponse

	booksData := s.repo.GetAllBook()

	for _, book := range booksData {
		booksResponse = append(booksResponse, utils.ToBookResponse(book))
	}

	return booksResponse
}
