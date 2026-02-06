package comment

import (
	"database/sql"

	"github.com/Tapan176/go-practice/internal"
	entityschema "github.com/Tapan176/go-practice/webApiV1/model/dto"
)

func GetAllCommentsService(dbClient *sql.DB, postID int) ([]entityschema.Comment, error) {
	comments, err := GetCommentsByPostID(dbClient, postID)
	if err != nil {
		return nil, internal.FetchCommentFailed
	}
	return comments, nil
}

func AddCommentService(dbClient *sql.DB, articleID, userID int, comment string) (*entityschema.Comment, error) {
	commentData, err := CreateCommentDal(dbClient, comment, userID, articleID)
	if err != nil {
		return nil, internal.InternalServerError
	}
	return commentData, nil
}

func UpdateCommentService(dbClient *sql.DB, commentID int, commentBody string) error {
	err := UpdateCommentDal(dbClient, commentID, commentBody)
	if err != nil {
		if err == sql.ErrNoRows {
			return internal.CommentNotFound
		}
		return internal.InternalServerError
	}
	return nil
}

func DeleteCommentService(dbClient *sql.DB, commentID int) error {
	err := DeleteCommentDal(dbClient, commentID)
	if err != nil {
		if err == sql.ErrNoRows {
			return internal.CommentNotFound
		}
		return internal.InternalServerError
	}
	return nil
}
