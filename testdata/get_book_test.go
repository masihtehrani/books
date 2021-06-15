package testdata_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/masihtehrani/books/app/ports/dto/dtogetbooks"

	"github.com/stretchr/testify/require"

	"github.com/masihtehrani/books/pkg/client"
)

func (t *BooksTests) TestGetBook() {
	ctx := context.Background()
	c := client.New(ctx, "", false)
	header := map[string][]string{
		"Content-Type": {"application/json"},
	}

	var response dtogetbooks.Response
	httplog, err := c.Request(ctx, fmt.Sprintf("http://%s:%s/books", ip, port), http.MethodGet, header, nil, &response, nil)
	require.NoError(t.Test, err)
	require.Equal(t.Test, httplog.Status, http.StatusOK)
	require.Equal(t.Test, len(response.Books), 1)
	require.Contains(t.Test, response.Books, "The History of a Crime")
	require.Equal(t.Test, response.Meta.Total, 1)

}
