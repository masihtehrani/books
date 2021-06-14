package http

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/entities/interfaces"
	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtodeletebook"
	"github.com/masihtehrani/books/pkg/server"
)

func deleteBookHandler(_ context.Context, iUseCases interfaces.IUseCases) server.MiddleFunc {
	return func(ctx context.Context, rawRequest server.RawRequest) (interface{}, error) {
		var request dtodeletebook.Request

		bookID, ok := rawRequest.Params["id"]
		if !ok || bookID == nil || bookID[0] == "" {
			return nil, fmt.Errorf("deleteBookHandler %w", structs.ErrParams)
		}

		request.BookID = bookID[0]

		userID, err := getUserID(rawRequest)
		if err != nil {
			return nil, fmt.Errorf("deleteBookHandler >> %w", err)
		}

		request.UserID = userID

		res, err := iUseCases.DeleteBook(ctx, &request)
		if err != nil {
			return nil, fmt.Errorf("deleteBookHandler >> %w", err)
		}

		return res, nil
	}
}
