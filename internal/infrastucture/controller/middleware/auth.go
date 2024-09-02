package middleware

import (
	"Kode_test/pkg/api/response"
	"context"
	"errors"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

var resourseIgnoreList = map[string]struct{}{
	"POST /auth/signup": {},
	"POST /auth/login":  {},
}

type JWT interface {
	ParseToken(tokenStr string) (int, error)
}

func ParseToken(jwt JWT) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			if _, ok := resourseIgnoreList[r.Method+" "+r.URL.Path]; !ok {

				//token, err := getTokenFromHeader(r)
				token := r.Header.Get("Authorization")
				if token == "" {
					render.JSON(w, r, response.Error("No token"))
					return
				}

				userID, err := jwt.ParseToken(token)
				if err != nil {
					render.JSON(w, r, response.Error("failed to parse token"))
					return
				}

				ctx = context.WithValue(r.Context(), "userID", userID)

			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header not set")
	}

	bearerTokenParts := strings.Split(authHeader, "Bearer")
	if len(bearerTokenParts) < 2 {
		return "", errors.New("authorization header has wrong format")
	}

	token := strings.TrimSpace(bearerTokenParts[1])
	if token == "" {
		return "", errors.New("authorization token not set")
	}

	return token, nil
}
