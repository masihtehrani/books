package server

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/masihtehrani/books/pkg/customeerror"
	"github.com/masihtehrani/books/pkg/logger"
)

func routeHandler(ctx context.Context, routes []*Router, logger *logger.Logger, jwtKey string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		path := req.URL.Path

		if path[len(path)-1:] != "/" {
			path += "/"
		}

		for _, route := range routes {
			if route.pattern.MatchString(path) && route.Method == req.Method {
				if route.IsNeedAuthenticate {
					err := checkJWT(ctx, res, req, jwtKey)
					if err != nil {
						makeResponseError(res, logger, err, http.StatusUnauthorized)

						return
					}
				}

				rawReq := makeRequest(res, req, logger)

				for i, param := range route.params {
					rawReq.Params[param] = []string{strings.Split(strings.TrimPrefix(path, route.rootPath), "/")[i]}
				}

				response, err := callFunction(ctx, route.FN, rawReq)
				if err != nil {
					makeResponseError(res, logger, err, http.StatusBadRequest)

					return
				}

				makeResponse(res, response, logger)

				return
			}
		}

		routeNotFoundErr := customeerror.New("100", "requested route is not exist")

		makeResponseError(res, logger, routeNotFoundErr, http.StatusNotFound)
	}
}

func (r *Router) makeParamsAndPth() error {
	pattern := r.Path
	r.rootPath = r.Path

	if strings.Contains(r.Path, ":") {
		pattern = r.Path[:strings.Index(r.Path, ":")]
		r.rootPath = r.Path[:strings.Index(r.Path, ":")]
	}

	params := make([]string, 0)

	for _, s := range strings.Split(r.Path, "/") {
		if strings.Contains(s, ":") {
			params = append(params, s[1:])
			pattern += `.*/`
		}
	}

	pathPattern, err := regexp.Compile(pattern)
	if err != nil {
		return fmt.Errorf("server pkg makeParamsAndPth regexpCompile >> %w", err)
	}

	r.pattern = pathPattern
	r.params = params

	return nil
}
