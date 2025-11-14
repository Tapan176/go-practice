package requestSchema

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterRequest struct {
	FirstName       string `json:"firstName" validate:"required,min=3,max=30,alpha"`
	LastName        string `json:"lastName" validate:"required,min=3,max=30,alpha"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=30"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordRequest struct {
	NewPassword     string `json:"newPassword" validate:"required,min=6,max=30"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=NewPassword"`
}
