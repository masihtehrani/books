package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/pkg/customeerror"
	"github.com/masihtehrani/books/pkg/logger"
)

const (
	filter = "filter"
	page   = "page"
	sort   = "sort"

	pageSizeDefault   = 10
	pageNumberDefault = 1
)

type MiddleFunc func(ctx context.Context, r RawRequest) (interface{}, error)

type RawRequest struct {
	Req    []byte
	Params map[string][]string
	Header map[string][]string
	structs.Query
}

func makeResponse(res http.ResponseWriter, response interface{}, logger *logger.Logger) {
	rBytes, err := json.Marshal(response)
	if err != nil {
		logger.Error.Println("HTTP", "HTTPRequestHandler", "Marshal response for call ", "error :", err)

		makeResponseError(res, logger, err, http.StatusInternalServerError)
	}

	res.WriteHeader(http.StatusOK)

	_, err = res.Write(rBytes)
	if err != nil {
		logger.Error.Println("HTTP", "HTTPRequestHandler for write to response error :", err)
	}
}

func makeResponseError(res http.ResponseWriter, logger *logger.Logger, err error, status int) {
	logger.Error.Println("HTTP", "HTTPRequestHandler", "error :", err)

	rBytes, err := json.Marshal(customeerror.NewErrorResponse(err, status))
	if err != nil {
		panic(err)
	}

	res.WriteHeader(status)

	_, err = res.Write(rBytes)
	if err != nil {
		panic(err)
	}
}

func callFunction(ctx context.Context, fn MiddleFunc, rawReq RawRequest) (interface{}, error) {
	response, err := fn(ctx, rawReq)

	return response, err
}

// nolint : gocritic
func makeRequest(res http.ResponseWriter, req *http.Request, logger *logger.Logger) RawRequest {
	rawReq := RawRequest{
		Req:    nil,
		Params: make(map[string][]string),
		Header: make(map[string][]string),
		Query: structs.Query{
			Filter: make(map[string]map[string][]string),
			Sort:   make(map[string][]string),
			Page: structs.Page{
				Number: pageNumberDefault,
				Size:   pageSizeDefault,
			},
		},
	}
	// fill server header.
	for k, values := range req.Header {
		rawReq.Header[k] = values
	}

	// fill makeRequest params & filters.
	val := req.URL.Query()
	for k, v := range val {
		if strings.HasPrefix(k, filter) {
			m := make(map[string][]string)
			m["eq"] = sliceToLower(v)
			rawReq.Filter[getKey(k, filter)] = m
		} else if strings.HasPrefix(k, page) {
			key := getKey(k, page)
			i, err := strconv.Atoi(v[0])
			if err == nil {
				if key == "number" {
					rawReq.Page.Number = i
				}
				if key == "size" {
					rawReq.Page.Size = i
				}
			}
		} else if strings.HasPrefix(k, sort) {
			rawReq.Sort[sliceToLower(v)[0]] = append(rawReq.Sort[sliceToLower(v)[0]], getKey(k, sort))
		} else {
			rawReq.Params[k] = v
		}
	}

	// fill makeRequest body.
	rByte, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logger.Error.Println("HTTP", "HTTPRequestHandler", "read Body makeRequest error: ", err)

		makeResponseError(res, logger, err, http.StatusBadRequest)
	}

	defer req.Body.Close()

	rawReq.Req = make([]byte, len(rByte))
	copy(rawReq.Req, rByte)

	return rawReq
}

func sliceToLower(s []string) []string {
	for i := range s {
		s[i] = strings.ToLower(s[i])
	}

	return s
}

func getKey(field, prefix string) string {
	reg := regexp.MustCompile(fmt.Sprintf("$1%s|\\[(.*?)\\]", prefix))
	res := reg.FindString(field)

	if res == "" {
		return res
	}

	res = strings.Trim(res, "[")
	res = strings.Trim(res, "]")

	return res
}
