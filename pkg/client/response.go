package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) makeResponse(_ context.Context, resp *http.Response, result, resultErr interface{}) ([]byte, error) {
	defer resp.Body.Close()

	p := result

	respBodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("makeResponse >> ioutil.ReadAll >> %w", err)
	}

	err = json.Unmarshal(respBodyByte, &result)
	if err != nil {
		return nil, fmt.Errorf("makeResonse json.Unmarshal() to result >>> errror: %w", err)
	}

	if result == p {
		err = json.Unmarshal(respBodyByte, &resultErr)
		if err != nil {
			return nil, fmt.Errorf("makeResponse >> json.Unmarshal() for resultErr >>> %w", err)
		}
	}

	return respBodyByte, nil
}
