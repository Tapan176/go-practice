package middleware

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/Tapan176/go-practice/constants"
)

func DatabaseMiddleware(db *sql.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), constants.DBContextKey, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetDB(r *http.Request) *sql.DB {
	db, ok := r.Context().Value(constants.DBContextKey).(*sql.DB)
	if !ok {
		return nil
	}
	return db
}
