package comment

import (
	"database/sql"
	"net/http"

	"github.com/Tapan176/go-practice/middleware"
)

func CommentRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	dbMiddleware := middleware.DatabaseMiddleware(db)

	router.Handle("GET /:articleId", middleware.Chain(
		middleware.ValidationMiddleware(GetAllCommentsValidation()),
		dbMiddleware,
		http.HandlerFunc(GetAllCommentsController),
	))
	router.Handle("POST /:articleId", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.ValidationMiddleware(AddCommentValidation()),
		dbMiddleware,
		http.HandlerFunc(CreateCommentController),
	))

	router.Handle("PUT /:articleId/:commentId", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.ValidationMiddleware(EditCommentValidation()),
		dbMiddleware,
		http.HandlerFunc(UpdateCommentController),
	))
	router.Handle("DELETE /:articleId/:commentId", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.ValidationMiddleware(DeleteCommentValidation()),
		dbMiddleware,
		http.HandlerFunc(DeleteCommentController),
	))

	return router
}
