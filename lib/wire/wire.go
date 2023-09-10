//go:build wireinject
// +build wireinject

package wire

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"

	"latihan_sqlc/interface/api/controller"
	"latihan_sqlc/internal/repositories"
	"latihan_sqlc/internal/service"
)

func InitializeApp(ctx context.Context, validate *validator.Validate, db *sql.DB) *controller.BookControllerImpl {
	wire.Build(
		repositories.NewBookRepositoriesImpl,
		service.NewBookServiceImpl,
		controller.NewBookControllerImpl,
	)
	return new(controller.BookControllerImpl)
}
