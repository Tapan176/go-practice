package user

import (
	"github.com/Tapan176/go-practice/middleware"
	requestSchema "github.com/Tapan176/go-practice/webApiV1/model/requestSchema"
)

func CreateUserValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: &requestSchema.CreateUserRequest{},
	}
}

func EditUserValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: &requestSchema.EditUserRequest{},
	}
}

func ChangePasswordValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: &requestSchema.ChangePasswordRequest{},
	}
}

func ChangeNameValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: &requestSchema.ChangeNameRequest{},
	}
}

func GetUserDetailsByEmailValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Body: &requestSchema.GetUserByEmailRequest{},
	}
}

func DeleteUserValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: &requestSchema.UserIDParam{},
	}
}

func GetUserDetailsByIdValidation() middleware.ValidationSchema {
	return middleware.ValidationSchema{
		Params: &requestSchema.UserIDParam{},
	}
}
