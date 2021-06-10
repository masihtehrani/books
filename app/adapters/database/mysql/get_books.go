package mysql

import (
	"context"
	"fmt"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtogetbooks"
	"github.com/masihtehrani/books/pkg/query"
)

func (m *Mysql) GetBooks(ctx context.Context, qu structs.Query) (*dtogetbooks.Response, error) {
	var response dtogetbooks.Response

	response.Books = make([]structs.Book, 0)

	q, err := query.New(m.db)
	if err != nil {
		return nil, fmt.Errorf("GetBooks >> query.New() >> %w", err)
	}

	//nolint: gofumpt
	q.Select = []string{"books.id", "books.title", "books.description", "books.is_published", "books.created_at",
		"authors.full_name", "authors.pseudonym"}
	q.Body = "from books inner join authors  on books.author_id = authors.id"
	q.QueryFilters = qu

	meta, err := q.Exec(ctx, &response.Books)
	if err != nil {
		return nil, fmt.Errorf("GetBooks >> q.Exec() >> %w", err)
	}

	response.Meta = meta

	return &response, nil
}
