package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/masihtehrani/books/app/entities/structs"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func checkJWT(_ context.Context, response http.ResponseWriter, request *http.Request, jwtSecretKey string) error {
	reqToken := request.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")

	if len(splitToken) != 2 { //nolint: gomnd
		response.WriteHeader(http.StatusUnauthorized)

		return structs.ErrUnauthorized
	}

	tknStr := strings.TrimSpace(splitToken[1])

	claims := &Claims{}

	jwtKey := []byte(jwtSecretKey)

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if !tkn.Valid {
		response.WriteHeader(http.StatusUnauthorized)

		return structs.ErrUnauthorized
	}

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			response.WriteHeader(http.StatusUnauthorized)

			return structs.ErrUnauthorized
		}

		response.WriteHeader(http.StatusBadRequest)

		return structs.ErrUnauthorized
	}

	request.Header.Add("X-User", claims.UserID)

	return nil
}

func CreateToken(ctx context.Context, jwtSecretKey string, userID string) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Hour)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", expirationTime, fmt.Errorf("jwt CreateToken >> %w", err)
	}

	return tokenString, expirationTime, nil
}
