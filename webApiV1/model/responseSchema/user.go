package responseSchema

import "time"

type UserDetailResponse struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	Role       string    `json:"role"`
	IsVerified bool      `json:"isVerified"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type UserListResponse struct {
	Users      []UserDetailResponse `json:"users"`
	Total      int                  `json:"total"`
	Page       int                  `json:"page,omitempty"`
	Limit      int                  `json:"limit,omitempty"`
	TotalPages int                  `json:"totalPages,omitempty"`
}

type PasswordChangeResponse struct {
	Message string `json:"message"`
}

type NameChangeResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Message   string `json:"message"`
}
