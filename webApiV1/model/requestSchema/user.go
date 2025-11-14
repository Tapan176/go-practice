package requestSchema

type UserIDParam struct {
	UserID string `json:"userId" validate:"required"`
}

type CreateUserRequest struct {
	FirstName       string `json:"firstName" validate:"required,min=3,max=30,alpha"`
	LastName        string `json:"lastName" validate:"required,min=3,max=30,alpha"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=6,max=30"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
	IsVerified      bool   `json:"isVerified" validate:"required"`
	UserRole        string `json:"userRole" validate:"required,oneof=admin user author"`
}

type EditUserRequest struct {
	FirstName  string `json:"firstName" validate:"required,min=3,max=30,alpha"`
	LastName   string `json:"lastName" validate:"required,min=3,max=30,alpha"`
	Email      string `json:"email" validate:"required,email"`
	IsVerified bool   `json:"isVerified" validate:"required"`
	UserRole   string `json:"userRole" validate:"required,oneof=admin user author"`
}

type ChangePasswordRequest struct {
	OldPassword        string `json:"oldPassword" validate:"required,min=6,max=30"`
	NewPassword        string `json:"newPassword" validate:"required,min=6,max=30"`
	ConfirmNewPassword string `json:"confirmNewPassword" validate:"required,eqfield=NewPassword"`
}

type ChangeNameRequest struct {
	FirstName string `json:"firstName" validate:"required,min=3,max=30,alpha"`
	LastName  string `json:"lastName" validate:"required,min=3,max=30,alpha"`
}

type GetUserByEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type ChangeUserPasswordRequest struct {
	OldPassword        string `json:"oldPassword"`
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}
