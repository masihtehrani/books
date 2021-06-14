package usecases

import (
	"context"

	"github.com/masihtehrani/books/app/entities/interfaces"
)

type UseCases struct {
	database interfaces.IDatabase
}

func New(_ context.Context, iDatabase interfaces.IDatabase) interfaces.IUseCases {
	return &UseCases{
		database: iDatabase,
	}
}
