package http

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/entities/interfaces"
	"github.com/masihtehrani/books/app/ports/dto/dtogetbooks"
	"github.com/masihtehrani/books/pkg/server"
)

func getBooksHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, rawRequest server.RawRequest) (interface{}, error) {
		var request dtogetbooks.Request

		request.Query = rawRequest.Query

		res, err := iUseCases.GetBooks(ctx, &request)
		if err != nil {
			return nil, fmt.Errorf("getCardsHandler >> %w", err)
		}

		return res, nil
	}
}
