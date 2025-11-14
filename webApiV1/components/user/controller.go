package user

import (
	"net/http"
	"strconv"

	"github.com/Tapan176/go-practice/constants"
	"github.com/Tapan176/go-practice/internal"
	"github.com/Tapan176/go-practice/middleware"
	entityschema "github.com/Tapan176/go-practice/webApiV1/model/entitySchema"
	"github.com/Tapan176/go-practice/webApiV1/model/requestSchema"
)

func GetAllUsersController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	users, err := GetAllUsersService(dbClient)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, users)
}

func GetUserByIDController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	userID := request.PathValue("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		internal.HandleError(response, "invalid_request_body")
		return
	}

	user, err := GetUserByIDService(dbClient, id)
	if err != nil {
		internal.HandleError(response, err)
		return
	}
	internal.SendOKResponse(response, user)
}

func GetUserByEmailController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	email := request.URL.Query().Get("email")
	user, err := GetUserByEmailService(dbClient, email)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendOKResponse(response, user)
}

func CreateUserController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	req := request.Context().Value(constants.ValidatedBodyKey).(*entityschema.User)
	user, err := CreateUserService(dbClient, req.FirstName, req.LastName, req.Email, req.PasswordHash, req.Role, req.IsVerified)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendCreatedResponse(response, user)
}

func UpdateUserController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	userID := request.PathValue("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		internal.HandleError(response, "invalid_request_body")
		return
	}
	req := request.Context().Value(constants.ValidatedBodyKey).(*entityschema.User)
	error := UpdateUserService(dbClient, id, req.FirstName, req.LastName, req.Email, req.Role, req.IsVerified)
	if error != nil {
		internal.HandleError(response, error)
		return
	}

	internal.SendNoContentResponse(response)
}

func DeleteUserController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	userID := request.PathValue("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		internal.HandleError(response, "invalid_request_body")
		return
	}
	err = DeleteUserService(dbClient, id)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendNoContentResponse(response)
}

func ChangeUserPasswordController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	req := request.Context().Value(constants.ValidatedBodyKey).(*requestSchema.ChangeUserPasswordRequest)

	userID := request.URL.Query().Get("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		internal.HandleError(response, "invalid_request_body")
		return
	}

	error := ChangeUserPasswordService(dbClient, id, req.OldPassword, req.NewPassword, req.ConfirmNewPassword)
	if error != nil {
		internal.HandleError(response, error)
		return
	}

	internal.SendNoContentResponse(response)
}

func ChangeUserNameController(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)
	req := request.Context().Value(constants.ValidatedBodyKey).(*entityschema.User)
	err := ChangeUserNameService(dbClient, req.ID, req.FirstName, req.LastName)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendNoContentResponse(response)
}
