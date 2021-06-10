package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/masihtehrani/books/app/entities/interfaces"
	"github.com/masihtehrani/books/app/ports/dto/dtocreatebook"
	"github.com/masihtehrani/books/pkg/server"
)

func createBookHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, rawRequest server.RawRequest) (interface{}, error) {
		var request dtocreatebook.Request

		userID, err := getUserID(rawRequest)
		if err != nil {
			return nil, fmt.Errorf("createBookHandler >> %w", err)
		}

		request.UserID = userID

		err = json.Unmarshal(rawRequest.Req, &request)
		if err != nil {
			return nil, fmt.Errorf("createBookHandler >> %w", err)
		}

		res, err := iUseCases.CreateBook(ctx, &request)
		if err != nil {
			return nil, fmt.Errorf("createBookHandler >> %w", err)
		}

		return res, nil
	}
}
