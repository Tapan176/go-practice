package responseSchema

import "time"

type PostResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Author     string    `json:"author,omitempty"`
	AuthorID   string    `json:"authorId,omitempty"`
	Category   string    `json:"category,omitempty"`
	CategoryID string    `json:"categoryId,omitempty"`
	Likes      int       `json:"likes"`
	Dislikes   int       `json:"dislikes"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type PostListResponse struct {
	Posts      []PostResponse `json:"posts"`
	Total      int            `json:"total"`
	Page       int            `json:"page,omitempty"`
	Limit      int            `json:"limit,omitempty"`
	TotalPages int            `json:"totalPages,omitempty"`
}

type CategoryResponse struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type LikeResponse struct {
	Liked    bool `json:"liked,omitempty"`
	Disliked bool `json:"disliked,omitempty"`
	Likes    int  `json:"likes"`
	Dislikes int  `json:"dislikes"`
}

type RatingResponse struct {
	Rating        int     `json:"rating"`
	AverageRating float64 `json:"averageRating,omitempty"`
	TotalRatings  int     `json:"totalRatings,omitempty"`
}
