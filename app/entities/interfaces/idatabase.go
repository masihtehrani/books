package interfaces

import (
	"context"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtogetbooks"
)

type IDatabase interface {
	GetBooks(ctx context.Context, query structs.Query) (*dtogetbooks.Response, error)
	CreateBook(ctx context.Context, book structs.Book, userID string) error
	SignUp(ctx context.Context, author structs.Author) error
	SignIn(ctx context.Context, username, password string) (string, error)
	Close() error
}
