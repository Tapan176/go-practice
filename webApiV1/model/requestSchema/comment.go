package requestSchema

type CommentIDParam struct {
	CommentID string `json:"commentId" validate:"required"`
}

type AddCommentRequest struct {
	Comment       string `json:"comment" validate:"required"`
	UserIdByAdmin string `json:"userIdByAdmin" validate:"omitempty"`
}

type EditCommentRequest struct {
	Comment string `json:"comment" validate:"required"`
}

type CreateCommentRequest struct {
	Comment string `json:"comment"`
	UserID  int    `json:"userId"`
}

type UpdateCommentRequest struct {
	Comment string `json:"comment"`
}
