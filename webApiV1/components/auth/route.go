package auth

import (
	"database/sql"
	"net/http"

	"github.com/Tapan176/go-practice/middleware"
)

func AuthRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	dbMiddleware := middleware.DatabaseMiddleware(db)

	router.Handle("POST /login", middleware.Chain(
		middleware.ValidationMiddleware(LoginValidation()),
		dbMiddleware,
		http.HandlerFunc(Login),
	))

	router.Handle("POST /register", middleware.Chain(
		middleware.ValidationMiddleware(RegisterValidation()),
		dbMiddleware,
		http.HandlerFunc(Register),
	))

	router.Handle("POST /logout", middleware.Chain(
		middleware.AuthenticationMiddleware,
		http.HandlerFunc(Logout),
	))

	return router
}
