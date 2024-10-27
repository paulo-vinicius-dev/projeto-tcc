package middlewares

import (
	"context"
	"net/http"
	"strings"
	"tcc-project/src/auth"
)

type contextKey string

const userIDKey = contextKey("user_id")

// BlacklistedToken ...
var BlacklistedToken = make(map[string]struct{})

// JWTMiddleWare ...
func JWTMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if _, exists := BlacklistedToken[tokenString]; exists {
			http.Error(w, "Token is blacklisted", http.StatusUnauthorized)
			return
		}

		claims, err := auth.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims["user_id"])
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
