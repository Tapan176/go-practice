package requestSchema

type GetBlogsQuery struct {
	Limit string `json:"limit" validate:"omitempty,numeric"`
	Page  string `json:"page" validate:"omitempty,numeric"`
}

type BlogIDParam struct {
	BlogID string `json:"blogId" validate:"required"`
}

type CategoryIDParam struct {
	CategoryID string `json:"categoryId" validate:"required"`
}

type CreatePostRequest struct {
	Title         string `json:"title" validate:"required"`
	Body          string `json:"body" validate:"required"`
	Category      string `json:"category" validate:"required"`
	UserIdByAdmin string `json:"userIdByAdmin" validate:"omitempty"`
}

type UpdatePostRequest struct {
	Title    string `json:"title" validate:"required"`
	Body     string `json:"body" validate:"required"`
	Category string `json:"category" validate:"required"`
}

type SearchArticleRequest struct {
	SearchString string `json:"searchString" validate:"required"`
	Limit        int    `json:"limit" validate:"omitempty,min=1"`
	Page         int    `json:"page" validate:"omitempty,min=1"`
}
