package auth

import (
	"database/sql"

	"github.com/Tapan176/go-practice/internal"
	"github.com/Tapan176/go-practice/webApiV1/components/user"
	entityschema "github.com/Tapan176/go-practice/webApiV1/model/dto"
	"golang.org/x/crypto/bcrypt"
)

func LoginService(dbClient *sql.DB, email, password string) (*entityschema.User, error) {
	user, err := user.GetUserByEmail(dbClient, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, internal.UserNotFound
		}
		return nil, internal.InternalServerError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, internal.IncorrectPassword
	}

	return user, nil
}

func RegisterService(dbClient *sql.DB, firstName, lastName, email, password string) (*entityschema.User, error) {
	exists, err := user.CheckUserExists(dbClient, email)
	if err != nil {
		return nil, internal.InternalServerError
	}

	if exists {
		return nil, internal.UserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, internal.InternalServerError
	}

	user, err := user.CreateUser(dbClient, firstName, lastName, email, string(hashedPassword), "user", false)
	if err != nil {
		return nil, internal.FailedToRegisterUser
	}

	return user, nil
}
