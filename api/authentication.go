package api

import (
	"context"
	"net/http"
	"strings"
)

func JWTValidationMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract the token from the Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Authorization header missing", http.StatusUnauthorized)
            return
        }

        // Expect the format "Bearer <token>"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, "Authorization header format must be Bearer <token>", http.StatusUnauthorized)
            return
        }

        token := parts[1]

        // Validate the token and extract the username
        id, err := ValidateToken(token)
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        // Set the username in the request context
        ctx := context.WithValue(r.Context(), "user_id", id)
        // Call the next handler with the new context
        next.ServeHTTP(w, r.WithContext(ctx))
    }
}
