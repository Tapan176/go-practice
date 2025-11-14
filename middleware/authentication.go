package middleware

import (
	"context"
	"net/http"

	"github.com/Tapan176/go-practice/constants"
	"github.com/Tapan176/go-practice/internal"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("userId")
		if userId == "" {
			internal.HandleError(w, "unauthorized")
			return
		}

		userRole := r.Header.Get("userRole")
		if userRole == "" {
			userRole = "user" // Default role
		}

		ctx := context.WithValue(r.Context(), constants.UserIDContextKey, userId)
		ctx = context.WithValue(ctx, constants.UserRoleContextKey, userRole)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserID(r *http.Request) string {
	userId, ok := r.Context().Value(constants.UserIDContextKey).(string)
	if !ok {
		return ""
	}
	return userId
}

func GetUserRole(r *http.Request) string {
	userRole, ok := r.Context().Value(constants.UserRoleContextKey).(string)
	if !ok {
		return ""
	}
	return userRole
}
