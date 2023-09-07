package service_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"latihan_sqlc/internal/models/web"
	mockService "latihan_sqlc/internal/service/mock"
)

var (
	mockServiceBook = mockService.IBoookService{}

	createBookData = web.BookCreateForm{Title: "One piece", Author: "Eichiiro Oda"}

	updateBookData = web.BookUpdateForm{Id: 2, Title: "Life is Short", Author: "Farda Ayu"}

	bookResponseReturn = func(bookId int, bookTitle, bookAuthor string) web.BookResponse {
		return web.BookResponse{Id: bookId, Author: bookAuthor, Title: bookTitle}
	}

	booksResponses = []web.BookResponse{
		{Id: 1, Author: "Eichiiro Oda", Title: "One piece"},
		{Id: 2, Author: "Farda Ayu", Title: "Life is Short"},
	}
)

func TestAddNewBook(t *testing.T) {
	mockServiceBook.Mock.On("AddNewBook", createBookData).Return(bookResponseReturn(1, createBookData.Title, createBookData.Author))
	bookResponse := mockServiceBook.AddNewBook(createBookData)
	assert.Equal(t, web.BookResponse{Id: 1, Author: createBookData.Author, Title: createBookData.Title}, bookResponse)
}

func TestUpdateBook(t *testing.T) {
	mockServiceBook.Mock.On("UpdateBook", updateBookData).Return(bookResponseReturn(updateBookData.Id, updateBookData.Title, updateBookData.Author))
	bookResponse := mockServiceBook.UpdateBook(updateBookData)
	assert.Equal(t, web.BookResponse{Id: updateBookData.Id, Author: updateBookData.Author, Title: updateBookData.Title}, bookResponse)
}

func TestFindBookById(t *testing.T) {
	mockServiceBook.Mock.On("FindBookById", 2).Return(bookResponseReturn(updateBookData.Id, updateBookData.Title, updateBookData.Author))
	bookResponse := mockServiceBook.FindBookById(2)
	assert.Equal(t, bookResponseReturn(updateBookData.Id, updateBookData.Title, updateBookData.Author), bookResponse)
}

func TestDeleteBook(t *testing.T) {
	// skip test because this method nothing return
}

func TestGetAllBooks(t *testing.T) {
	mockServiceBook.Mock.On("GetAllBooks").Return(booksResponses)
	bookResponses := mockServiceBook.GetAllBooks()
	assert.Equal(t, booksResponses, bookResponses)
}
