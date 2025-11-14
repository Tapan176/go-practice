package responseSchema

import "time"

type CommentResponse struct {
	ID        int                    `json:"id"`
	ArticleID int                    `json:"articleId"`
	UserID    int                    `json:"userId"`
	Comment   string                 `json:"comment"`
	Likes     int                    `json:"likes"`
	Dislikes  int                    `json:"dislikes"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
	Replies   []CommentReplyResponse `json:"replies,omitempty"`
}

type CommentReplyResponse struct {
	ID        int       `json:"id"`
	CommentID int       `json:"commentId"`
	UserID    int       `json:"userId"`
	Comment   string    `json:"comment"`
	Likes     int       `json:"likes"`
	Dislikes  int       `json:"dislikes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CommentListResponse struct {
	Comments   []CommentResponse `json:"comments"`
	Total      int               `json:"total"`
	Page       int               `json:"page,omitempty"`
	Limit      int               `json:"limit,omitempty"`
	TotalPages int               `json:"totalPages,omitempty"`
}
