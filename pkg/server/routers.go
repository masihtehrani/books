package server

import (
	"context"
	"net/http"
	"regexp"

	"github.com/masihtehrani/books/pkg/logger"
)

type Router struct {
	Path               string
	Method             string
	FN                 MiddleFunc
	params             []string
	rootPath           string
	pattern            *regexp.Regexp
	IsNeedAuthenticate bool
}

func routers(ctx context.Context, mux *http.ServeMux, routers []*Router, logger *logger.Logger, jwtKey string) {
	for _, router := range routers {
		err := router.makeParamsAndPth()
		if err != nil {
			panic(err)
		}
	}

	mux.HandleFunc("/", routeHandler(ctx, routers, logger, jwtKey))
}
