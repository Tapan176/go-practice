package user

import (
	"database/sql"

	"github.com/Tapan176/go-practice/internal"
	entityschema "github.com/Tapan176/go-practice/webApiV1/model/entitySchema"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsersService(dbClient *sql.DB) ([]entityschema.User, error) {
	users, err := GetAllUsers(dbClient)
	if err != nil {
		return nil, internal.InternalServerError
	}
	return users, nil
}

func GetUserByIDService(dbClient *sql.DB, userID int) (*entityschema.User, error) {
	user, err := GetUserByID(dbClient, userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, internal.UserNotFound
		}
		return nil, internal.InternalServerError
	}
	return user, nil
}

func GetUserByEmailService(dbClient *sql.DB, email string) (*entityschema.User, error) {
	user, err := GetUserByEmail(dbClient, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, internal.UserNotFound
		}
		return nil, internal.InternalServerError
	}
	return user, nil
}

func CreateUserService(dbClient *sql.DB, firstName, lastName, email, passwordHash, role string, isVerified bool) (*entityschema.User, error) {
	existingUser, err := GetUserByEmail(dbClient, email)
	if err == nil && existingUser != nil {
		return nil, internal.UserAlreadyExists
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, internal.InternalServerError
	}
	user, err := CreateUser(dbClient, firstName, lastName, email, string(hashedPassword), role, isVerified)
	if err != nil {
		return nil, internal.InternalServerError
	}
	return user, nil
}

func UpdateUserService(dbClient *sql.DB, userID int, firstName, lastName, email, role string, isVerified bool) error {
	err := UpdateUser(dbClient, firstName, lastName, email, role, userID, isVerified)
	if err != nil {
		return internal.InternalServerError
	}
	return nil
}

func DeleteUserService(dbClient *sql.DB, userID int) error {
	err := DeleteUser(dbClient, userID)
	if err != nil {
		return internal.InternalServerError
	}
	return nil
}

func ChangeUserPasswordService(dbClient *sql.DB, userID int, oldPassword, newPassword, confirmNewPassword string) error {
	user, err := GetUserByID(dbClient, userID)
	if err != nil || user == nil {
		return internal.UserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return internal.IncorrectPassword
	}

	if newPassword != confirmNewPassword {
		return internal.PasswordAndConfirmPasswordDoNotMatch
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(newPassword)) == nil {
		return internal.NewPasswordSameAsOldPassword
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return internal.InternalServerError
	}

	err = UpdateUserPassword(dbClient, userID, string(hashedPassword))
	if err != nil {
		return internal.InternalServerError
	}
	return nil
}

func ChangeUserNameService(dbClient *sql.DB, userID int, firstName, lastName string) error {
	err := UpdateUserName(dbClient, userID, firstName, lastName)
	if err != nil {
		return internal.InternalServerError
	}
	return nil
}
