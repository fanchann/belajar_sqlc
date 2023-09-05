package main

import (
	"fmt"

	"latihan_sqlc/config"
	"latihan_sqlc/internal/models"
	"latihan_sqlc/internal/repositories"
)

var (
	data = []models.Books{
		{Author: "Farda Ayu", Title: "Ketika Senja Perlahan tenggelam"},
		{Author: "John Doe", Title: "Hello World"},
	}
)

func main() {
	cfg := config.New("./.env")
	db := config.MysqlConnection(cfg)

	// repository
	repo := repositories.NewBookRepositoriesImpl(db)

	//repo.AddBook(data[1])

	fmt.Printf("repo.GetAllBook(): %v\n", repo.GetAllBook())
}
