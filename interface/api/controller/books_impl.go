package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"latihan_sqlc/internal/models/web"
	"latihan_sqlc/internal/service"
	"latihan_sqlc/utils"
)

type BookControllerImpl struct {
	Service *service.BookServiceImpl
}

func NewBookControllerImpl(Service *service.BookServiceImpl) *BookControllerImpl {
	return &BookControllerImpl{Service: Service}
}

func (cntrl *BookControllerImpl) AddNewBook(c *fiber.Ctx) error {
	addForm := new(web.BookCreateForm)
	errBodyParser := c.BodyParser(addForm)
	utils.PanicIfError(errBodyParser)

	bookResponse := cntrl.Service.AddNewBook(*addForm)

	response := web.WebResponse{Code: 200, Status: "OK", Data: bookResponse}

	return c.JSON(response)
}

func (cntrl *BookControllerImpl) UpdateBook(c *fiber.Ctx) error {
	bookId := c.Params("bookId")
	updateForm := new(web.BookUpdateForm)
	errBodyParser := c.BodyParser(updateForm)
	utils.PanicIfError(errBodyParser)

	id, err := strconv.Atoi(bookId)
	utils.PanicIfError(err)

	updateForm.Id = id

	bookResponse := cntrl.Service.UpdateBook(*updateForm)

	response := web.WebResponse{Code: 200, Status: "OK", Data: bookResponse}
	return c.JSON(response)
}

func (cntrl *BookControllerImpl) Delete(c *fiber.Ctx) error {
	bookId := c.Params("bookId")

	id, err := strconv.Atoi(bookId)
	utils.PanicIfError(err)

	cntrl.Service.Delete(id)
	response := web.WebResponse{Code: 200, Status: "OK"}
	return c.JSON(response)
}

func (cntrl *BookControllerImpl) FindBookById(c *fiber.Ctx) error {
	bookId := c.Params("bookId")

	id, err := strconv.Atoi(bookId)
	utils.PanicIfError(err)
	bookResponse := cntrl.Service.FindBookById(id)
	response := web.WebResponse{Code: 200, Status: "OK", Data: bookResponse}
	return c.JSON(response)
}

func (cntrl *BookControllerImpl) GetAllBooks(c *fiber.Ctx) error {
	c.Response().Header.Add("Cache-Time", "6000")
	bookResponses := cntrl.Service.GetAllBooks()
	response := web.WebResponse{Code: 200, Status: "OK", Data: bookResponses}
	return c.JSON(response)
}
