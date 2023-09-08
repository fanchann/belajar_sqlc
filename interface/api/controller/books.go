package controller

import "github.com/gofiber/fiber/v2"

type IBookController interface {
	AddNewBook(c fiber.Ctx) error
	UpdateBook(c fiber.Ctx) error
	Delete(c fiber.Ctx) error
	FindBookById(c fiber.Ctx) error
	GetAllBooks(c fiber.Ctx) error
}
