package utils

import (
	"latihan_sqlc/internal/models/domain"
	"latihan_sqlc/internal/models/web"
)

func ToBookResponse(domain domain.Books) web.BookResponse {
	return web.BookResponse{
		Id:     domain.Id,
		Title:  domain.Title,
		Author: domain.Author,
	}
}
