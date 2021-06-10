package server

import (
	"errors"
	"net/http"

	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/pkg/logger"
)

func recoveryMiddleware(h http.Handler, logger *logger.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t) // nolint: goerr113
				case error:
					err = t
				default:
					err = structs.ErrUnknown
				}
				logger.Error.Printf("panic: %s", err.Error())
				makeResponseError(w, logger, err, http.StatusInternalServerError)
			}
		}()
		h.ServeHTTP(w, r)
	})
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
