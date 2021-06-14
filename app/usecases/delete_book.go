package usecases

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/ports/dto/dtodeletebook"
)

func (u *UseCases) DeleteBook(ctx context.Context, request *dtodeletebook.Request) (*dtodeletebook.Response, error) {
	err := u.database.DeleteBook(ctx, request.UserID, request.BookID)
	if err != nil {
		return nil, fmt.Errorf("usecases.DeleteBook >> %w", err)
	}

	return &dtodeletebook.Response{Success: true}, nil
}
