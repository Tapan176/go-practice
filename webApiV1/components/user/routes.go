package user

import (
	"database/sql"
	"net/http"

	"github.com/Tapan176/go-practice/middleware"
)

func UserRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	dbMiddleware := middleware.DatabaseMiddleware(db)

	router.Handle("GET /", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.AuthorizationMiddleware,
		dbMiddleware,
		http.HandlerFunc(GetAllUsersController),
	))
	router.Handle("POST /", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.AuthorizationMiddleware,
		middleware.ValidationMiddleware(CreateUserValidation()),
		dbMiddleware,
		http.HandlerFunc(CreateUserController),
	))

	router.Handle("GET /email", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.AuthorizationMiddleware,
		middleware.ValidationMiddleware(GetUserDetailsByEmailValidation()),
		dbMiddleware,
		http.HandlerFunc(GetUserByEmailController),
	))

	router.Handle("GET /{id}", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.AuthorizationMiddleware,
		middleware.ValidationMiddleware(GetUserDetailsByIdValidation()),
		dbMiddleware,
		http.HandlerFunc(GetUserByIDController),
	))
	router.Handle("PUT /{id}", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.AuthorizationMiddleware,
		middleware.ValidationMiddleware(EditUserValidation()),
		dbMiddleware,
		http.HandlerFunc(UpdateUserController),
	))
	router.Handle("DELETE /{id}", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.AuthorizationMiddleware,
		middleware.ValidationMiddleware(DeleteUserValidation()),
		dbMiddleware,
		http.HandlerFunc(DeleteUserController),
	))

	router.Handle("PUT /change-password", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.ValidationMiddleware(ChangePasswordValidation()),
		dbMiddleware,
		http.HandlerFunc(ChangeUserPasswordController),
	))
	router.Handle("PUT /change-name", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.ValidationMiddleware(ChangeNameValidation()),
		dbMiddleware,
		http.HandlerFunc(ChangeUserNameController),
	))

	return router
}
