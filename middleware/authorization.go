package middleware

import (
	"net/http"

	"github.com/Tapan176/go-practice/internal"
)

func AuthorizationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := GetUserRole(r)

		if userRole != "admin" {
			internal.HandleError(w, "unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func RequireRole(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := GetUserRole(r)

			hasRole := false
			for _, role := range allowedRoles {
				if userRole == role {
					hasRole = true
					break
				}
			}

			if !hasRole {
				internal.HandleError(w, "unauthorized")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
