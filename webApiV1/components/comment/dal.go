package comment

import (
	"database/sql"

	entityschema "github.com/Tapan176/go-practice/webApiV1/model/dto"
)

func GetCommentsByPostID(dbClient *sql.DB, postID int) ([]entityschema.Comment, error) {
	query := `SELECT * FROM "comments" WHERE "articleId" = $1`
	rows, err := dbClient.Query(query, postID)
	if err != nil {
		return []entityschema.Comment{}, err
	}
	defer rows.Close()
	var comments []entityschema.Comment
	for rows.Next() {
		var c entityschema.Comment
		if err := rows.Scan(&c.ID, &c.Comment, &c.UserId, &c.ArticleId, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return []entityschema.Comment{}, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func CreateCommentDal(dbClient *sql.DB, comment string, userID, articleID int) (*entityschema.Comment, error) {
	query := `
		INSERT INTO "comments" ("comment", "userId", "articleId", "createdAt", "updatedAt")
		VALUES ($1, $2, $3, NOW(), NOW())
		RETURNING "id", "comment", "userId", "articleId", "createdAt", "updatedAt"
	`
	var insertedComment entityschema.Comment
	err := dbClient.QueryRow(query, comment, userID, articleID).Scan(
		&insertedComment.ID, &insertedComment.Comment, &insertedComment.UserId, &insertedComment.ArticleId,
		&insertedComment.CreatedAt, &insertedComment.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &insertedComment, nil
}

func UpdateCommentDal(dbClient *sql.DB, commentID int, commentBody string) error {
	query := `
		UPDATE "comments"
		SET "comment" = $1, "updatedAt" = NOW()
		WHERE "id" = $2
	`
	_, err := dbClient.Exec(query, commentBody, commentID)
	return err
}

func DeleteCommentDal(dbClient *sql.DB, commentID int) error {
	query := `DELETE FROM "comments" WHERE "id" = $1`
	_, err := dbClient.Exec(query, commentID)
	return err
}

func GetUserIdByCommentID(dbClient *sql.DB, commentID int) (int, error) {
	query := `SELECT "userId" FROM "comments" WHERE "id" = $1`
	var userID int
	err := dbClient.QueryRow(query, commentID).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
