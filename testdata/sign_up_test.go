package testdata_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"

	"github.com/masihtehrani/books/app/ports/dto/dtosignup"
	"github.com/masihtehrani/books/pkg/client"
)

func (t *BooksTests) TestSignUp() {
	ctx := context.Background()
	c := client.New(ctx, "", false)
	header := map[string][]string{
		"Content-Type": {"application/json"},
	}
	body := dtosignup.Request{
		FullName:  "Victor Hugo",
		Pseudonym: "Hugo",
		Username:  "VictorHugo",
		Password:  "VictorHugo",
	}
	var response dtosignup.Response
	httplog, err := c.Request(ctx, fmt.Sprintf("http://%s:%s/sign-up", ip, port), http.MethodPost, header, body, &response, nil)
	require.NoError(t.Test, err)
	require.Equal(t.Test, httplog.Status, http.StatusOK)
	require.Equal(t.Test, response.Success, true)

}
