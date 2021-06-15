package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) makeRequest(ctx context.Context, url, method string, header map[string][]string,
	params interface{}) (*http.Request, error) {
	var body io.Reader

	if params != nil {
		b, err := c.makeRequestBody(ctx, params)
		if err != nil {
			return nil, err
		}

		body = b
	}

	req, err := http.NewRequestWithContext(ctx, checkMethods(ctx, method), url, body)
	if err != nil {
		return nil, fmt.Errorf("makeRequest >> server.NewRequestWithContext >> %w", err)
	}

	c.makeRequestHeader(ctx, req, header)

	return req, nil
}

func checkMethods(_ context.Context, method string) string {
	switch strings.ToUpper(method) {
	case http.MethodConnect:
		return http.MethodConnect
	case http.MethodDelete:
		return http.MethodDelete
	case http.MethodHead:
		return http.MethodHead
	case http.MethodOptions:
		return http.MethodOptions
	case http.MethodPatch:
		return http.MethodPatch
	case http.MethodPost:
		return http.MethodPost
	case http.MethodPut:
		return http.MethodPut
	case http.MethodTrace:
		return http.MethodTrace
	default:
		return http.MethodGet
	}
}

func (c *Client) makeRequestBody(_ context.Context, params interface{}) (io.Reader, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("makeRequestBody >> json.Marshal >> %w", err)
	}

	return bytes.NewBuffer(body), nil
}

func (c *Client) makeRequestHeader(_ context.Context, req *http.Request, header map[string][]string) {
	for k, v := range header {
		req.Header[http.CanonicalHeaderKey(k)] = v
	}
}
