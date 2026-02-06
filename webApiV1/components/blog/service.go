package blog

import (
	"database/sql"
	"strconv"

	"github.com/Tapan176/go-practice/internal"
	entityschema "github.com/Tapan176/go-practice/webApiV1/model/dto"
)

func GetPostsService(dbClient *sql.DB, limit, page int) ([]entityschema.Post, error) {
	if limit == 0 {
		limit = 5
	}
	if page == 0 {
		page = 1
	}

	posts, err := GetPostsDal(dbClient, limit, page)
	if err != nil {
		return nil, internal.FetchBlogsFailed
	}

	return posts, nil
}

func CreatePostService(dbClient *sql.DB, title, body, userId, category string) (*entityschema.Post, error) {
	categoryData, err := GetCategoryByTitle(dbClient, category)
	if err != nil {
		if err == sql.ErrNoRows {
			categoryData, err = CreateCategory(dbClient, category)
			if err != nil {
				return nil, internal.AddBlogFailed
			}
		} else {
			return nil, internal.AddBlogFailed
		}
	}
	post, err := CreatePostDal(dbClient, title, body, userId, strconv.Itoa(categoryData.ID))
	if err != nil {
		return nil, internal.AddBlogFailed
	}

	return post, nil
}

func UpdatePostService(dbClient *sql.DB, postID, title, body, category string) (*entityschema.Post, error) {
	var categoryID string

	if category != "" {
		categoryData, err := GetCategoryByTitle(dbClient, category)
		if err != nil {
			if err == sql.ErrNoRows {
				categoryData, err = CreateCategory(dbClient, category)
				if err != nil {
					return nil, internal.UpdateBlogFailed
				}
			} else {
				return nil, internal.UpdateBlogFailed
			}
		}
		categoryID = strconv.Itoa(categoryData.ID)
	}

	post, err := UpdatePostDal(dbClient, postID, title, body, categoryID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, internal.BlogNotFound
		}
		return nil, internal.UpdateBlogFailed
	}

	return post, nil
}

func DeletePostService(dbClient *sql.DB, postID string) error {
	err := DeletePostDal(dbClient, postID)
	if err != nil {
		return internal.DeleteBlogFailed
	}
	return nil
}

func SearchArticleService(dbClient *sql.DB, searchString string, limit, page int) ([]entityschema.Post, error) {
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	posts, err := SearchArticleDal(dbClient, searchString, limit, page)
	if err != nil {
		return nil, internal.FetchBlogsFailed
	}

	return posts, nil
}

func GetArticleByIDService(dbClient *sql.DB, blogID string) (*entityschema.Post, error) {
	post, err := GetArticleByIDDal(dbClient, blogID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, internal.BlogNotFound
		}
		return nil, internal.FetchBlogsFailed
	}
	return post, nil
}

func GetUserArticlesService(dbClient *sql.DB, userID string) ([]entityschema.Post, error) {
	posts, err := GetUserArticlesDal(dbClient, userID)
	if err != nil {
		return nil, internal.FetchBlogsFailed
	}
	return posts, nil
}

func LikeArticleService(dbClient *sql.DB, blogID, userID string) (map[string]interface{}, error) {
	alreadyLiked, err := CheckUserLiked(dbClient, blogID, userID)
	if err != nil {
		return nil, err
	}

	if alreadyLiked {
		err = RemoveLike(dbClient, blogID, userID)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"message": "Like removed successfully",
			"liked":   false,
		}, nil
	}

	alreadyDisliked, err := CheckUserDisliked(dbClient, blogID, userID)
	if err != nil {
		return nil, err
	}
	if alreadyDisliked {
		err = RemoveDislike(dbClient, blogID, userID)
		if err != nil {
			return nil, err
		}
	}

	err = AddLike(dbClient, blogID, userID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "Article liked successfully",
		"liked":   true,
	}, nil
}

func DislikeArticleService(dbClient *sql.DB, blogID, userID string) (map[string]interface{}, error) {
	alreadyDisliked, err := CheckUserDisliked(dbClient, blogID, userID)
	if err != nil {
		return nil, err
	}

	if alreadyDisliked {
		err = RemoveDislike(dbClient, blogID, userID)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"message":  "Dislike removed successfully",
			"disliked": false,
		}, nil
	}

	alreadyLiked, err := CheckUserLiked(dbClient, blogID, userID)
	if err != nil {
		return nil, err
	}
	if alreadyLiked {
		err = RemoveLike(dbClient, blogID, userID)
		if err != nil {
			return nil, err
		}
	}

	err = AddDislike(dbClient, blogID, userID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message":  "Article disliked successfully",
		"disliked": true,
	}, nil
}

func RateAuthorService(dbClient *sql.DB, blogID, userID string, rating int) (map[string]interface{}, error) {
	article, err := GetArticleByIDDal(dbClient, blogID)
	if err != nil {
		return nil, err
	}

	authorID := article.UserId

	existingRating, err := GetUserRatingForAuthor(dbClient, authorID, userID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err != sql.ErrNoRows {
		err = UpdateAuthorRating(dbClient, existingRating.ID, rating)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"message": "Rating updated successfully",
		}, nil
	}

	err = AddAuthorRating(dbClient, authorID, userID, rating)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "Author rated successfully",
	}, nil
}

func GetArticlesByCategoryService(dbClient *sql.DB, categoryID string, limit, page int) ([]entityschema.Post, error) {
	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	posts, err := GetArticlesByCategoryDal(dbClient, categoryID, limit, page)
	if err != nil {
		return nil, internal.FetchBlogsFailed
	}
	return posts, nil
}
