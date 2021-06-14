package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/masihtehrani/books/app/entities/interfaces"
	"github.com/masihtehrani/books/app/entities/structs"
	"github.com/masihtehrani/books/app/ports/dto/dtoversion"
	"github.com/masihtehrani/books/pkg/logger"
	"github.com/masihtehrani/books/pkg/server"
)

const (
	version = "/version"
	books   = "/books"
	book    = "/books/:id"
	signIn  = "/sign-in"
	signUp  = "/sign-up"
)

func New(ctx context.Context, iUseCases interfaces.IUseCases, ver dtoversion.Response,
	logger *logger.Logger) (*http.Server, error) {
	ip := os.Getenv("HTTP_IP")
	portEnv := os.Getenv("HTTP_PORT")
	jwtKey := os.Getenv("JWT_SECRET_KEY")

	port, err := strconv.ParseUint(portEnv, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("New >> strconv.ParseUint >> %w", err)
	}

	if ip == "" || port == 0 {
		return nil, structs.ErrEmptyIPANDPORT
	}

	if jwtKey == "" {
		return nil, structs.ErrJwtSecretToken
	}

	routers := []*server.Router{
		{Path: signUp, Method: http.MethodPost, FN: signUpHandler(ctx, iUseCases), IsNeedAuthenticate: false},
		{Path: signIn, Method: http.MethodPost, FN: signInHandler(ctx, iUseCases, jwtKey), IsNeedAuthenticate: false},
		{Path: books, Method: http.MethodGet, FN: getBooksHandler(ctx, iUseCases), IsNeedAuthenticate: false},
		{Path: books, Method: http.MethodPost, FN: createBookHandler(ctx, iUseCases), IsNeedAuthenticate: true},
		{Path: book, Method: http.MethodPut, FN: updateBookHandler(ctx, iUseCases), IsNeedAuthenticate: true},
		{Path: book, Method: http.MethodDelete, FN: deleteBookHandler(ctx, iUseCases), IsNeedAuthenticate: true},
		{Path: version, Method: http.MethodGet, FN: versionHandler(ctx, ver), IsNeedAuthenticate: false},
	}

	serverHTTP, err := server.New(ctx, jwtKey, ip, port, routers, logger)
	if err != nil {
		return nil, fmt.Errorf("New >> server.New >> %w", err)
	}

	return serverHTTP, nil
}

func getUserID(request server.RawRequest) (string, error) {
	userID, ok := request.Header[structs.XUser]
	if !ok || userID == nil || userID[0] == "" {
		return "", fmt.Errorf("getUserID %w", structs.ErrUserID)
	}

	return userID[0], nil
}
