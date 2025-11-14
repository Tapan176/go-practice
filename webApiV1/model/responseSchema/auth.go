package responseSchema

import "time"

type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token,omitempty"`
}

type RegisterResponse struct {
	User    UserResponse `json:"user"`
	Message string       `json:"message"`
}

type ForgotPasswordResponse struct {
	Message string `json:"message"`
	Email   string `json:"email"`
}

type ResetPasswordResponse struct {
	Message string `json:"message"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}
