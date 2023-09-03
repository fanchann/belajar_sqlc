package repositories

import (
	"context"
	"database/sql"
	"log"

	"latihan_sqlc/internal/models"
	"latihan_sqlc/internal/mysql/sqlc"
)

var ctx = context.Background()

type BookRepositoriesImpl struct {
	Db *sql.DB
}

func NewBookRepositoriesImpl(db *sql.DB) IBoookRepositories {
	return &BookRepositoriesImpl{Db: db}
}

func (r *BookRepositoriesImpl) AddBook(book models.Books) models.Books {

	sqlcQueries := sqlc.New(r.Db)

	result, err := sqlcQueries.AddNewBook(ctx, sqlc.AddNewBookParams{Author: book.Author, Title: book.Title})
	if err != nil {
		panic(err)
	}
	lastInsertId, _ := result.LastInsertId()
	book.Id = int(lastInsertId)

	return book
}

func (r *BookRepositoriesImpl) GetBookById(id int) (models.Books, error) {

	sqlcQueries := sqlc.New(r.Db)
	result, err := sqlcQueries.GetBookById(ctx, int32(id))
	if err != nil {
		return models.Books{}, err
	}
	return models.Books{Id: int(result.ID), Author: result.Author, Title: result.Title}, nil
}

func (r *BookRepositoriesImpl) UpdateBookById(book models.Books) models.Books {

	sqlcQueries := sqlc.New(r.Db)
	err := sqlcQueries.UpdateBook(ctx, sqlc.UpdateBookParams{ID: int32(book.Id), Author: book.Author})
	if err != nil {
		log.Fatalf(err.Error())
		return models.Books{}
	}
	return book
}

func (r *BookRepositoriesImpl) GetAllBook() []models.Books {

	var booksModels []models.Books
	sqlcQueries := sqlc.New(r.Db)
	books, err := sqlcQueries.GetAllBooks(ctx)

	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, book := range books {
		booksModels = append(booksModels, sqlcBookToModel(book))
	}
	return booksModels
}

func (r *BookRepositoriesImpl) DeleteBookById(id int) error {

	sqlcQueries := sqlc.New(r.Db)
	err := sqlcQueries.DeleteAuthor(ctx, int32(id))
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}

func sqlcBookToModel(book sqlc.Book) models.Books {
	return models.Books{Id: int(book.ID), Title: book.Title, Author: book.Author}
}
