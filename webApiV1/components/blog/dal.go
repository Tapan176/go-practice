package blog

import (
	"database/sql"

	entityschema "github.com/Tapan176/go-practice/webApiV1/model/dto"
)

func GetPostsDal(dbClient *sql.DB, limit, page int) ([]entityschema.Post, error) {
	query := `SELECT * FROM "articles" LIMIT $1 OFFSET $2`
	rows, err := dbClient.Query(query, limit, (page-1)*limit)
	if err != nil {
		return []entityschema.Post{}, err
	}
	defer rows.Close()

	var posts []entityschema.Post
	for rows.Next() {
		var p entityschema.Post
		if err := rows.Scan(&p); err != nil {
			return []entityschema.Post{}, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}

func CreatePostDal(dbClient *sql.DB, title, body, userID, categoryId string) (*entityschema.Post, error) {
	query := `
		INSERT INTO "articles" ("title", "body", "userId", "categoryId", "createdAt", "updatedAt")
		VALUES ($1, $2, $3, $4, NOW(), NOW())
		RETURNING "id", "title", "body", "userId", "categoryId", "createdAt", "updatedAt", "likes", "dislikes"
	`

	var post entityschema.Post
	err := dbClient.QueryRow(query, title, body, userID, categoryId).Scan(
		&post.ID, &post.Title, &post.Body, &post.UserId, &post.CategoryId,
		&post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes,
	)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetCategoryByTitle(dbClient *sql.DB, title string) (*entityschema.Category, error) {
	query := `SELECT "id", "title" FROM "categories" WHERE "title" = $1`

	var category entityschema.Category
	err := dbClient.QueryRow(query, title).Scan(&category.ID, &category.Title)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func CreateCategory(dbClient *sql.DB, title string) (*entityschema.Category, error) {
	query := `
		INSERT INTO "categories" ("title", "createdAt", "updatedAt")
		VALUES ($1, NOW(), NOW())
		RETURNING "id", "title"
	`

	var category entityschema.Category
	err := dbClient.QueryRow(query, title).Scan(&category.ID, &category.Title)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func UpdatePostDal(dbClient *sql.DB, postID, title, body, categoryID string) (*entityschema.Post, error) {
	query := `
		UPDATE "articles"
		SET "title" = COALESCE(NULLIF($2, ''), "title"),
		    "body" = COALESCE(NULLIF($3, ''), "body"),
		    "categoryId" = COALESCE(NULLIF($4, ''), "categoryId"),
		    "updatedAt" = NOW()
		WHERE "id" = $1
		RETURNING "id", "title", "body", "userId", "categoryId", "createdAt", "updatedAt", "likes", "dislikes"
	`

	var post entityschema.Post
	err := dbClient.QueryRow(query, postID, title, body, categoryID).Scan(
		&post.ID, &post.Title, &post.Body, &post.UserId, &post.CategoryId,
		&post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes,
	)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func DeletePostDal(dbClient *sql.DB, postID string) error {
	query := `DELETE FROM "articles" WHERE "id" = $1`

	result, err := dbClient.Exec(query, postID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func SearchArticleDal(dbClient *sql.DB, searchString string, limit, page int) ([]entityschema.Post, error) {
	query := `
		SELECT "id", "title", "body", "userId", "categoryId", "createdAt", "updatedAt", "likes", "dislikes"
		FROM "articles"
		WHERE "title" ILIKE $1 OR "body" ILIKE $1
		ORDER BY "createdAt" DESC
		LIMIT $2 OFFSET $3
	`

	searchPattern := "%" + searchString + "%"
	rows, err := dbClient.Query(query, searchPattern, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entityschema.Post
	for rows.Next() {
		var post entityschema.Post
		err := rows.Scan(
			&post.ID, &post.Title, &post.Body, &post.UserId, &post.CategoryId,
			&post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func GetArticleByIDDal(dbClient *sql.DB, blogID string) (*entityschema.Post, error) {
	query := `
		SELECT "id", "title", "body", "userId", "categoryId", "createdAt", "updatedAt", "likes", "dislikes"
		FROM "articles"
		WHERE "id" = $1
	`

	var post entityschema.Post
	err := dbClient.QueryRow(query, blogID).Scan(
		&post.ID, &post.Title, &post.Body, &post.UserId, &post.CategoryId,
		&post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes,
	)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetUserArticlesDal(dbClient *sql.DB, userID string) ([]entityschema.Post, error) {
	query := `
		SELECT "id", "title", "body", "userId", "categoryId", "createdAt", "updatedAt", "likes", "dislikes"
		FROM "articles"
		WHERE "userId" = $1
		ORDER BY "createdAt" DESC
	`

	rows, err := dbClient.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entityschema.Post
	for rows.Next() {
		var post entityschema.Post
		err := rows.Scan(
			&post.ID, &post.Title, &post.Body, &post.UserId, &post.CategoryId,
			&post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func CheckUserLiked(dbClient *sql.DB, blogID, userID string) (bool, error) {
	query := `SELECT COUNT(*) FROM "likedBlog" WHERE "blogId" = $1 AND "userId" = $2`

	var count int
	err := dbClient.QueryRow(query, blogID, userID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CheckUserDisliked(dbClient *sql.DB, blogID, userID string) (bool, error) {
	query := `SELECT COUNT(*) FROM "dislikedBlog" WHERE "blogId" = $1 AND "userId" = $2`

	var count int
	err := dbClient.QueryRow(query, blogID, userID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func AddLike(dbClient *sql.DB, blogID, userID string) error {
	tx, err := dbClient.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`INSERT INTO "likedBlog" ("blogId", "userId", "createdAt") VALUES ($1, $2, NOW())`, blogID, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE "articles" SET "likes" = "likes" + 1 WHERE "id" = $1`, blogID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func RemoveLike(dbClient *sql.DB, blogID, userID string) error {
	tx, err := dbClient.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM "likedBlog" WHERE "blogId" = $1 AND "userId" = $2`, blogID, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE "articles" SET "likes" = GREATEST("likes" - 1, 0) WHERE "id" = $1`, blogID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func AddDislike(dbClient *sql.DB, blogID, userID string) error {
	tx, err := dbClient.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`INSERT INTO "dislikedBlog" ("blogId", "userId", "createdAt") VALUES ($1, $2, NOW())`, blogID, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE "articles" SET "dislikes" = "dislikes" + 1 WHERE "id" = $1`, blogID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func RemoveDislike(dbClient *sql.DB, blogID, userID string) error {
	tx, err := dbClient.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`DELETE FROM "dislikedBlog" WHERE "blogId" = $1 AND "userId" = $2`, blogID, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(`UPDATE "articles" SET "dislikes" = GREATEST("dislikes" - 1, 0) WHERE "id" = $1`, blogID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func GetUserRatingForAuthor(dbClient *sql.DB, authorID, userID string) (*entityschema.AuthorRating, error) {
	query := `SELECT "id", "authorId", "userId", "rating" FROM "authorRating" WHERE "authorId" = $1 AND "userId" = $2`

	var rating entityschema.AuthorRating
	err := dbClient.QueryRow(query, authorID, userID).Scan(&rating.ID, &rating.AuthorID, &rating.UserID, &rating.Rating)
	if err != nil {
		return nil, err
	}

	return &rating, nil
}

func AddAuthorRating(dbClient *sql.DB, authorID, userID string, rating int) error {
	query := `INSERT INTO "authorRating" ("authorId", "userId", "rating", "createdAt") VALUES ($1, $2, $3, NOW())`

	_, err := dbClient.Exec(query, authorID, userID, rating)
	return err
}

func UpdateAuthorRating(dbClient *sql.DB, ratingID int, newRating int) error {
	query := `UPDATE "authorRating" SET "rating" = $1, "updatedAt" = NOW() WHERE "id" = $2`

	_, err := dbClient.Exec(query, newRating, ratingID)
	return err
}

func GetArticlesByCategoryDal(dbClient *sql.DB, categoryID string, limit, page int) ([]entityschema.Post, error) {
	query := `
		SELECT "id", "title", "body", "userId", "categoryId", "createdAt", "updatedAt", "likes", "dislikes"
		FROM "articles"
		WHERE "categoryId" = $1
		ORDER BY "createdAt" DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := dbClient.Query(query, categoryID, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entityschema.Post
	for rows.Next() {
		var post entityschema.Post
		err := rows.Scan(
			&post.ID, &post.Title, &post.Body, &post.UserId, &post.CategoryId,
			&post.CreatedAt, &post.UpdatedAt, &post.Likes, &post.Dislikes,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
