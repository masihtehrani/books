package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/pkg/logger"
)

func New(ctx context.Context, jwtKey, ip string, port uint64, router []*Router, logger *logger.Logger,
	isTest bool) (*http.Server,
	error) {
	if ip == "" || port == 0 || router == nil || len(router) == 0 {
		return nil, structs.ErrEmptyIPANDPORT
	}

	logger.Info.Printf("server on IP: %s & Port: %d", ip, port)

	mux := http.NewServeMux()

	routers(ctx, mux, router, logger, jwtKey)

	HTTPServer := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: recoveryMiddleware(jsonMiddleware(mux), logger),
	}

	if isTest {
		go func() {
			err := HTTPServer.ListenAndServe()
			if err != nil {
				logger.Error.Fatalf("server New server error :%s", HTTPServer.ListenAndServe())
			}
		}()
	} else {
		err := HTTPServer.ListenAndServe()
		if err != nil {
			return nil, fmt.Errorf("server New server error :%w", HTTPServer.ListenAndServe())
		}
	}

	return HTTPServer, nil
}
