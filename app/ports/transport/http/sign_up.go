package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/masihtehrani/books/app/entities/interfaces"
	"github.com/masihtehrani/books/app/ports/dto/dtosignup"
	"github.com/masihtehrani/books/pkg/server"
)

func signUpHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, rawRequest server.RawRequest) (interface{}, error) {
		var request dtosignup.Request

		err := json.Unmarshal(rawRequest.Req, &request)
		if err != nil {
			return nil, fmt.Errorf("signUpHandler >> %w", err)
		}

		res, err := iUseCases.SignUp(ctx, &request)
		if err != nil {
			return nil, fmt.Errorf("signUpHandler >> %w", err)
		}

		return res, nil
	}
}
