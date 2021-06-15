package testdata_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtocreatebook"

	"github.com/stretchr/testify/require"

	"github.com/masihtehrani/books/pkg/client"
)

func (t *BooksTests) TestCreateBook() {
	ctx := context.Background()
	c := client.New(ctx, "", false)
	header := map[string][]string{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", token)},
	}
	body := dtocreatebook.Request{
		Book: structs.Book{
			Title:       "The History of a Crime",
			Description: "The History of a Crime (French: Histoire d'un crime, 1877) is an essay by Victor Hugo about Napoleon III's takeover of France. External links[edit].",
			IsPublished: true,
		},
	}
	var response dtocreatebook.Response
	httplog, err := c.Request(ctx, fmt.Sprintf("http://%s:%s/books", ip, port), http.MethodPost, header, body, &response, nil)
	require.NoError(t.Test, err)
	require.Equal(t.Test, httplog.Status, http.StatusOK)
	require.Equal(t.Test, response.Success, true)

}
