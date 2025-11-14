package routes

import (
	"database/sql"
	"net/http"

	"github.com/Tapan176/go-practice/webApiV1/components/auth"
	"github.com/Tapan176/go-practice/webApiV1/components/blog"
	"github.com/Tapan176/go-practice/webApiV1/components/comment"
	"github.com/Tapan176/go-practice/webApiV1/components/user"
)

func IndexRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("/auth/", http.StripPrefix("/auth", auth.AuthRouter(db)))
	router.Handle("/blog/", http.StripPrefix("/blog", blog.BlogRouter(db)))
	router.Handle("/comment/", http.StripPrefix("/comment", comment.CommentRouter(db)))
	router.Handle("/user/", http.StripPrefix("/user", user.UserRouter(db)))

	return router
}
