package blog

import (
	"database/sql"
	"net/http"

	"github.com/Tapan176/go-practice/middleware"
)

func BlogRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	dbMiddleware := middleware.DatabaseMiddleware(db)

	router.Handle("GET /", middleware.Chain(
		middleware.ValidationMiddleware(GetArticlesValidation()),
		dbMiddleware,
		http.HandlerFunc(GetPostsHandler),
	))

	router.Handle("POST /", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.ValidationMiddleware(AddArticleValidation()),
		dbMiddleware,
		http.HandlerFunc(CreatePostHandler),
	))

	router.Handle("GET /search", middleware.Chain(
		middleware.ValidationMiddleware(SearchArticleValidation()),
		dbMiddleware,
		http.HandlerFunc(SearchArticleHandler),
	))

	router.Handle("GET /myblogs", middleware.Chain(
		middleware.AuthenticationMiddleware,
		dbMiddleware,
		http.HandlerFunc(GetUserArticlesHandler),
	))

	router.Handle("GET /{blogId}", middleware.Chain(
		middleware.ValidationMiddleware(GetArticlesByIDValidation()),
		dbMiddleware,
		http.HandlerFunc(GetArticleByIDHandler),
	))

	router.Handle("PUT /{blogId}", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.ValidationMiddleware(EditArticleValidation()),
		dbMiddleware,
		http.HandlerFunc(UpdatePostHandler),
	))

	router.Handle("DELETE /{blogId}", middleware.Chain(
		middleware.AuthenticationMiddleware,
		middleware.ValidationMiddleware(DeleteArticleValidation()),
		dbMiddleware,
		http.HandlerFunc(DeletePostHandler),
	))

	router.Handle("PUT /{blogId}/like", middleware.Chain(
		middleware.AuthenticationMiddleware,
		dbMiddleware,
		http.HandlerFunc(LikeArticleHandler),
	))

	router.Handle("PUT /{blogId}/dislike", middleware.Chain(
		middleware.AuthenticationMiddleware,
		dbMiddleware,
		http.HandlerFunc(DislikeArticleHandler),
	))

	router.Handle("POST /{blogId}/rate-author", middleware.Chain(
		middleware.AuthenticationMiddleware,
		dbMiddleware,
		http.HandlerFunc(RateAuthorHandler),
	))

	router.Handle("GET /categories/{categoryId}/blogs", middleware.Chain(
		middleware.ValidationMiddleware(GetArticlesByCategoryValidation()),
		dbMiddleware,
		http.HandlerFunc(GetArticlesByCategoryHandler),
	))

	return router
}
