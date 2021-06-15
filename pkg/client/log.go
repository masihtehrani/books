package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) httpLog(header map[string][]string, params interface{}, req *http.Request, url string,
	response *http.Response, resp []byte) (string, string) {
	paramsByte, _ := json.Marshal(params)

	reqResult := fmt.Sprintf(
		"%s %s %s\n"+ // POST /api/v1/users HTTP/1.1
			"%s\n"+ // header
			"%s\n\n\n\n", // request
		req.Method, url, req.Proto,
		c.map2string(header),
		string(paramsByte),
	)

	respResult := fmt.Sprintf(
		"%s %s\n"+ // HTTP/1.1 200 OK
			"%s\n"+ // response header
			"%s", // response
		response.Proto, response.Status,
		c.map2string(response.Header),
		string(resp),
	)

	return reqResult, respResult
}

func (c *Client) map2string(m map[string][]string) string {
	var s string

	for k, hs := range m {
		s += k + ": "

		for _, h := range hs {
			s += h + ","
		}

		s += "\n"
	}

	return s
}
