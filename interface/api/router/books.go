package router

import (
	"context"
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/utils"

	"latihan_sqlc/config"
	"latihan_sqlc/exception"
	"latihan_sqlc/internal/models/web"
	"latihan_sqlc/lib/wire"
)

func InitializeRouter(appConfig *string) *fiber.App {
	ctx := context.Background()
	validate := validator.New()
	db := config.MysqlConnection(config.New(*appConfig))
	initializeApp := wire.InitializeApp(ctx, validate, db)

	c := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {

			code := fiber.StatusInternalServerError

			var e *exception.HttpError
			if errors.As(err, &e) {
				code = int(e.StatusCode())
			}

			var valErr *exception.ValidationError
			if errors.As(err, &valErr) {
				ctx.Status(fiber.ErrBadRequest.Code)
				return ctx.JSON(web.NewWebResponse(fiber.ErrBadGateway.Code, err.Error(), nil))
			}

			var notFound *exception.NotFoundErr
			if errors.As(err, &notFound) {
				ctx.Status(fiber.ErrNotFound.Code)
				return ctx.JSON(web.NewWebResponse(fiber.ErrNotFound.Code, "data not found", nil))
			}

			return ctx.Status(code).JSON(web.NewWebResponse(fiber.ErrBadGateway.Code, err.Error(), nil))
		}})

	c.Use(logger.New(logger.Config{
		Format:     "[${ip}]:${port}|${status}|${method}|${path}|${latency}\n",
		TimeFormat: "02-Jan-2006",
	}))
	c.Use(cache.New(cache.Config{
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.Path())
		},
	}))

	c.Use(recover.New())

	c.Get("/api/books", initializeApp.GetAllBooks)
	c.Get("/api/book/:bookId", initializeApp.FindBookById)

	c.Post("/api/book", initializeApp.AddNewBook)

	c.Put("/api/book/:bookId", initializeApp.UpdateBook)

	c.Delete("/api/book/:bookId", initializeApp.Delete)

	return c
}
