package middleware

import (
	"context"
	"net/http"
	"projeto_chat_backend/pkg/auth"
	"strings"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		claims, err := auth.ParseToken(bearerToken[1])
		if err != nil {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "userID", claims.UserID))
		next.ServeHTTP(w, r)
	}
}
