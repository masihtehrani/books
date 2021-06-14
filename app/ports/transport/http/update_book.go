package http

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/masihtehrani/books/app/entities/interfaces"
	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtoupdatebook"
	"github.com/masihtehrani/books/pkg/server"
)

func updateBookHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, rawRequest server.RawRequest) (interface{}, error) {
		var request dtoupdatebook.Request

		bookID, ok := rawRequest.Params["id"]
		if !ok || bookID == nil || bookID[0] == "" {
			return nil, fmt.Errorf("updateBookHandler %w", structs.ErrParams)
		}

		request.BookID = bookID[0]

		userID, err := getUserID(rawRequest)
		if err != nil {
			return nil, fmt.Errorf("updateBookHandler >> %w", err)
		}

		request.UserID = userID

		err = json.Unmarshal(rawRequest.Req, &request)
		if err != nil {
			return nil, fmt.Errorf("createBookHandler >> %w", err)
		}

		res, err := iUseCases.UpdateBook(ctx, &request)
		if err != nil {
			return nil, fmt.Errorf("updateBookHandler >> %w", err)
		}

		return res, nil
	}
}
