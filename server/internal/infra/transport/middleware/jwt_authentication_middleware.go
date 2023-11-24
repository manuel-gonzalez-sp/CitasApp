package middleware

import (
	"citasapp/internal/infra/utils"
	"context"
	"net/http"
	"strings"
)

func JWTAuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 {
			http.Error(w, "Authorization header is malformed", http.StatusUnauthorized)
			return
		}

		tokenString := headerParts[1]
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Add claims to the context so the handlers can use the user information
		ctx := context.WithValue(r.Context(), utils.UserClaimsContextKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
