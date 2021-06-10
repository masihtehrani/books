package usecases

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/ports/dto/dtogetbooks"
)

func (u *UseCases) GetBooks(ctx context.Context, request *dtogetbooks.Request) (*dtogetbooks.Response, error) {
	res, err := u.database.GetBooks(ctx, request.Query)
	if err != nil {
		return nil, fmt.Errorf("usecases.GetBooks >> %w", err)
	}

	return res, nil
}
