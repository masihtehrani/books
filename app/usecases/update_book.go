package usecases

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtoupdatebook"
)

func (u *UseCases) UpdateBook(ctx context.Context, request *dtoupdatebook.Request) (*dtoupdatebook.Response, error) {
	if request.Title == "" || request.Description == "" {
		return nil, structs.ErrValidationCreateBook
	}

	err := u.database.UpdateBook(ctx, request.Book, request.UserID, request.BookID)
	if err != nil {
		return nil, fmt.Errorf("usecases.UpdateBook >> %w", err)
	}

	return &dtoupdatebook.Response{Success: true}, nil
}
