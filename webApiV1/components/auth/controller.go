package auth

import (
	"encoding/gob"
	"net/http"

	"github.com/Tapan176/go-practice/constants"
	"github.com/Tapan176/go-practice/internal"
	"github.com/Tapan176/go-practice/middleware"
	entityschema "github.com/Tapan176/go-practice/webApiV1/model/dto"
	requestSchema "github.com/Tapan176/go-practice/webApiV1/model/requestSchema"
	"github.com/gorilla/sessions"
)

var cookieData = sessions.NewCookieStore([]byte("super-secret-key"))

func init() {
	gob.Register(entityschema.SessionUser{})

	cookieData.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
}

func Login(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	req := request.Context().Value(constants.ValidatedBodyKey).(*requestSchema.LoginRequest)

	user, err := LoginService(dbClient, req.Email, req.Password)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	session, _ := cookieData.Get(request, "session")
	sessionUser := entityschema.SessionUser{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
	}
	session.Values["user"] = sessionUser

	if err := session.Save(request, response); err != nil {
		internal.HandleError(response, "internal_server_error")
		return
	}

	internal.SendOKResponse(response, user)
}

func Register(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	dbClient := middleware.GetDB(request)

	req := request.Context().Value(constants.ValidatedBodyKey).(*requestSchema.RegisterRequest)

	user, err := RegisterService(dbClient, req.FirstName, req.LastName, req.Email, req.Password)
	if err != nil {
		internal.HandleError(response, err)
		return
	}

	internal.SendCreatedResponse(response, user)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	session, err := cookieData.Get(request, "session")
	if err != nil {
		internal.HandleError(response, "internal_server_error")
		return
	}

	if user := session.Values["user"]; user != nil {
		delete(session.Values, "user")
		session.Save(request, response)
		internal.SendNoContentResponse(response)
		return
	}

	internal.HandleError(response, "user_has_not_login")
}
