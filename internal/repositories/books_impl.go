package repositories

import (
	"context"
	"database/sql"
	"log"

	"latihan_sqlc/internal/models/domain"
	"latihan_sqlc/internal/mysql/sqlc"
)

type BookRepositoriesImpl struct {
	Ctx context.Context
	Db  *sql.DB
}

func NewBookRepositoriesImpl(ctx context.Context, db *sql.DB) IBoookRepositories {
	return &BookRepositoriesImpl{Db: db, Ctx: ctx}
}

func (r *BookRepositoriesImpl) AddBook(book domain.Books) domain.Books {

	sqlcQueries := sqlc.New(r.Db)

	result, err := sqlcQueries.AddNewBook(r.Ctx, sqlc.AddNewBookParams{Author: book.Author, Title: book.Title})
	if err != nil {
		panic(err)
	}
	lastInsertId, _ := result.LastInsertId()
	book.Id = int(lastInsertId)

	return book
}

func (r *BookRepositoriesImpl) GetBookById(id int) (domain.Books, error) {

	sqlcQueries := sqlc.New(r.Db)
	result, err := sqlcQueries.GetBookById(r.Ctx, int32(id))
	if err != nil {
		return domain.Books{}, err
	}
	return domain.Books{Id: int(result.ID), Author: result.Author, Title: result.Title}, nil
}

func (r *BookRepositoriesImpl) UpdateBookById(book domain.Books) domain.Books {

	sqlcQueries := sqlc.New(r.Db)
	err := sqlcQueries.UpdateBook(r.Ctx, sqlc.UpdateBookParams{ID: int32(book.Id), Author: book.Author})
	if err != nil {
		log.Fatalf(err.Error())
		return domain.Books{}
	}
	return book
}

func (r *BookRepositoriesImpl) GetAllBook() []domain.Books {

	var booksdomain []domain.Books
	sqlcQueries := sqlc.New(r.Db)
	books, err := sqlcQueries.GetAllBooks(r.Ctx)

	if err != nil {
		log.Fatalf(err.Error())
	}

	for _, book := range books {
		booksdomain = append(booksdomain, sqlcBookToModel(book))
	}
	return booksdomain
}

func (r *BookRepositoriesImpl) DeleteBookById(book domain.Books) error {

	sqlcQueries := sqlc.New(r.Db)
	err := sqlcQueries.DeleteAuthor(r.Ctx, int32(book.Id))
	if err != nil {
		log.Fatalf(err.Error())
		return err
	}
	return nil
}

func sqlcBookToModel(book sqlc.Book) domain.Books {
	return domain.Books{Id: int(book.ID), Title: book.Title, Author: book.Author}
}
