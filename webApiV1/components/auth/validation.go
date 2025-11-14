package auth

import (
	"github.com/Tapan176/go-practice/middleware"
	requestSchema "github.com/Tapan176/go-practice/webApiV1/model/requestSchema"
)

func LoginValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: requestSchema.LoginRequest{},
	}
}

func RegisterValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: requestSchema.RegisterRequest{},
	}
}

func ForgotPasswordValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: requestSchema.ForgotPasswordRequest{},
	}
}

func ResetPasswordValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: requestSchema.ResetPasswordRequest{},
	}
}
