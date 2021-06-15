package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/masihtehrani/books/app/entities/structs"
)

type Client struct {
	client *http.Client
}

type HTTPLog struct {
	Request  string
	Response string
	URL      string
	Status   int
	Duration time.Duration
}

func createHTTPClient(_ context.Context, proxy string, isSkipSSL bool) *http.Client {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: isSkipSSL} //nolint: gosec

	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}

	client := http.DefaultClient
	client.Transport = transport

	return client
}

func New(ctx context.Context, proxy string, isDisabledSSL bool) *Client {
	return &Client{client: createHTTPClient(ctx, proxy, isDisabledSSL)}
}

func (c *Client) Request(ctx context.Context, url, method string, header map[string][]string,
	params, result, resultErr interface{}) (*HTTPLog, error) {
	if url == "" {
		return nil, structs.ErrClientURLNotFilled
	}

	req, err := c.makeRequest(ctx, url, method, header, params)
	if err != nil {
		return nil, err
	}

	start := time.Now()

	response, err := c.do(ctx, req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	duration := time.Since(start)

	resp, err := c.makeResponse(ctx, response, &result, &resultErr)
	if err != nil {
		return nil, err
	}

	reqResult, respResult := c.httpLog(header, params, req, url, response, resp)

	httpLog := HTTPLog{
		Request:  reqResult,
		Response: respResult,
		URL:      url,
		Status:   response.StatusCode,
		Duration: duration,
	}

	return &httpLog, nil
}

func (c *Client) do(_ context.Context, req *http.Request) (*http.Response, error) {
	response, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do >> c.client.Do >> %w", err)
	}

	return response, nil
}
