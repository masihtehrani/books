package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/masihtehrani/books/app/entities/interfaces"
	"github.com/masihtehrani/books/app/ports/dto/dtosignin"
	"github.com/masihtehrani/books/pkg/server"
)

func signInHandler(_ context.Context, iUseCases interfaces.IUseCases, jwtKey string) server.MiddleFunc {
	return func(ctx context.Context, rawRequest server.RawRequest) (interface{}, error) {
		var request dtosignin.Request

		err := json.Unmarshal(rawRequest.Req, &request)
		if err != nil {
			return nil, fmt.Errorf("signInHandler >> %w", err)
		}

		userID, err := iUseCases.SignIn(ctx, &request)
		if err != nil {
			return nil, fmt.Errorf("signInHandler >> %w", err)
		}

		token, expireTime, err := server.CreateToken(ctx, jwtKey, userID)
		if err != nil {
			return nil, fmt.Errorf("signInHandler >> %w", err)
		}

		return &dtosignin.Response{
			Token:    token,
			ExpireAt: expireTime,
		}, nil
	}
}
