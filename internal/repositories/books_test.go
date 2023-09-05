package repositories_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"latihan_sqlc/internal/models"
	mockRepo "latihan_sqlc/internal/repositories/mock"
)

var (
	mockBookRepo = mockRepo.IBoookRepositories{mock.Mock{}}
	fakeData     = []models.Books{
		{Id: 1, Title: "Title 1", Author: "Author 1"},
		{Id: 2, Title: "Title 2", Author: "Author 2"},
		{Id: 3, Title: "Title 3", Author: "Author 3"},
		{Id: 4, Title: "Title 4", Author: "Author 4"},
	}
)

func TestAddBook(t *testing.T) {
	mockBookRepo.On("AddBook", fakeData[0]).Return(models.Books{Id: 1, Title: "Title 1", Author: "Author 1"})
	successInsert := mockBookRepo.AddBook(fakeData[0])
	assert.Equal(t, successInsert, models.Books{Id: 1, Title: "Title 1", Author: "Author 1"})
}

func TestUpdateBook(t *testing.T) {
	mockBookRepo.On("UpdateBookById", fakeData[1]).Return(fakeData[1])
	successUpdate := mockBookRepo.UpdateBookById(fakeData[1])
	assert.Equal(t, successUpdate, fakeData[1])
}

func TestDeleteBook(t *testing.T) {
	mockBookRepo.On("DeleteBookById", 1).Return(nil)
	err := mockBookRepo.DeleteBookById(1)
	assert.Nil(t, err)
}

func TestGetBookById(t *testing.T) {
	// success
	t.Run("Success", func(t *testing.T) {
		mockBookRepo.On("GetBookById", 1).Return(fakeData[0], nil)
		book, err := mockBookRepo.GetBookById(1)
		assert.Nil(t, err)
		assert.Equal(t, book, fakeData[0])
	})

	// error
	t.Run("Error", func(t *testing.T) {
		mockBookRepo.On("GetBookById", 100).Return(models.Books{}, errors.New("book not found"))
		book, err := mockBookRepo.GetBookById(100)
		assert.NotNil(t, err)
		assert.Equal(t, book, models.Books{})
	})
}

func TestGetAllBooks(t *testing.T) {
	mockBookRepo.On("GetAllBook").Return(fakeData)
	books := mockBookRepo.GetAllBook()
	assert.Equal(t, books, fakeData)
}
