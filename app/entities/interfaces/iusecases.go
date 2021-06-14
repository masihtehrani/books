package interfaces

import (
	"context"

	"github.com/masihtehrani/books/app/ports/dto/dtocreatebook"
	"github.com/masihtehrani/books/app/ports/dto/dtodeletebook"
	"github.com/masihtehrani/books/app/ports/dto/dtogetbooks"
	"github.com/masihtehrani/books/app/ports/dto/dtosignin"
	"github.com/masihtehrani/books/app/ports/dto/dtosignup"
	"github.com/masihtehrani/books/app/ports/dto/dtoupdatebook"
)

type IUseCases interface {
	GetBooks(ctx context.Context, request *dtogetbooks.Request) (*dtogetbooks.Response, error)
	CreateBook(ctx context.Context, request *dtocreatebook.Request) (*dtocreatebook.Response, error)
	UpdateBook(ctx context.Context, request *dtoupdatebook.Request) (*dtoupdatebook.Response, error)
	DeleteBook(ctx context.Context, request *dtodeletebook.Request) (*dtodeletebook.Response, error)
	SignUp(ctx context.Context, request *dtosignup.Request) (*dtosignup.Response, error)
	SignIn(ctx context.Context, request *dtosignin.Request) (string, error)
}
