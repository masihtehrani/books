package usecases

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtocreatebook"
)

func (u *UseCases) CreateBook(ctx context.Context, request *dtocreatebook.Request) (*dtocreatebook.Response, error) {
	if request.Title == "" || request.Description == "" {
		return nil, structs.ErrValidationCreateBook
	}

	err := u.database.CreateBook(ctx, request.Book, request.UserID)
	if err != nil {
		return nil, fmt.Errorf("usecases.CreateBook >> %w", err)
	}

	return &dtocreatebook.Response{Success: true}, nil
}
