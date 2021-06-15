package testdata_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"

	"github.com/masihtehrani/books/app/ports/dto/dtosignin"

	"github.com/masihtehrani/books/pkg/client"
)

func (t *BooksTests) TestSignIn() {
	ctx := context.Background()
	c := client.New(ctx, "", false)
	header := map[string][]string{
		"Content-Type": {"application/json"},
	}
	body := dtosignin.Request{
		Username: "VictorHugo",
		Password: "VictorHugo",
	}
	var response dtosignin.Response
	httplog, err := c.Request(ctx, fmt.Sprintf("http://%s:%s/sign-in", ip, port), http.MethodPost, header, body, &response, nil)
	require.NoError(t.Test, err)
	require.Equal(t.Test, httplog.Status, http.StatusOK)
	require.NotEmpty(t.Test, response.Token)
	token = response.Token
	//assert.Equal(t.Test, response.Success, true)

}
