package usecases

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtosignup"
)

func (u *UseCases) SignUp(ctx context.Context, request *dtosignup.Request) (*dtosignup.Response, error) {
	if request.Password == "" || request.Username == "" || request.Pseudonym == "" || request.FullName == "" {
		return nil, structs.ErrValidationSignUp
	}

	author := structs.Author{
		FullName:  request.FullName,
		Pseudonym: request.Pseudonym,
		Username:  request.Username,
		Password:  request.Password,
	}

	err := u.database.SignUp(ctx, author)
	if err != nil {
		return nil, fmt.Errorf("usecases.SignUp >> %w", err)
	}

	return &dtosignup.Response{Success: true}, nil
}
