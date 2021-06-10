package usecases

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtosignin"
)

func (u *UseCases) SignIn(ctx context.Context, request *dtosignin.Request) (string, error) {
	if request.Password == "" || request.Username == "" {
		return "", structs.ErrValidationSigIn
	}

	userID, err := u.database.SignIn(ctx, request.Username, request.Password)
	if err != nil {
		return "", fmt.Errorf("usecases.SignIn >> %w", err)
	}

	return userID, nil
}
