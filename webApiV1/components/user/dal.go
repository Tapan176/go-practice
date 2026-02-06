package user

import (
	"database/sql"

	entityschema "github.com/Tapan176/go-practice/webApiV1/model/dto"
)

func GetAllUsers(dbClient *sql.DB) ([]entityschema.User, error) {
	query := `SELECT * FROM "users"`
	rows, err := dbClient.Query(query)
	if err != nil {
		return []entityschema.User{}, err
	}
	defer rows.Close()

	var users []entityschema.User
	for rows.Next() {
		var u entityschema.User
		if err := rows.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName); err != nil {
			return []entityschema.User{}, err
		}
		users = append(users, u)
	}

	return users, nil
}

func CheckUserExists(dbClient *sql.DB, email string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)`

	err := dbClient.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func GetUserByID(dbClient *sql.DB, userID int) (*entityschema.User, error) {
	query := `SELECT "id", "email", "firstName", "lastName" FROM "users" WHERE "id" = $1`
	var user entityschema.User
	err := dbClient.QueryRow(query, userID).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(dbClient *sql.DB, email string) (*entityschema.User, error) {
	query := `SELECT "id", "email", "firstName", "lastName", "passwordHash", "role" FROM "users" WHERE "email" = $1`
	var user entityschema.User
	err := dbClient.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(dbClient *sql.DB, firstName, lastName, email, passwordHash, role string, isVerified bool) (*entityschema.User, error) {
	query := `
		INSERT INTO "users" ("firstName", "lastName", "email", "passwordHash", "isVerified", "role")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING "id", "email", "firstName", "lastName", "isVerified", "role"
	`
	var user entityschema.User
	err := dbClient.QueryRow(query, firstName, lastName, email, passwordHash, isVerified, role).Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.IsVerified, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(dbClient *sql.DB, firstName, lastName, email, role string, userID int, isVerified bool) error {
	query := `
		UPDATE "users"
		SET "firstName" = $1, "lastName" = $2, "email" = $3, "role" = $4, "isVerified" = $5, "updatedAt" = NOW()
		WHERE "id" = $6
	`
	_, err := dbClient.Exec(query, firstName, lastName, email, role, isVerified, userID)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(dbClient *sql.DB, userID int) error {
	query := `DELETE FROM "users" WHERE "id" = $1`
	_, err := dbClient.Exec(query, userID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserPassword(dbClient *sql.DB, userID int, newPasswordHash string) error {
	query := `
		UPDATE "users"
		SET "passwordHash" = $1, "updatedAt" = NOW()
		WHERE "id" = $2
	`
	_, err := dbClient.Exec(query, newPasswordHash, userID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUserName(dbClient *sql.DB, userID int, firstName, lastName string) error {
	query := `
		UPDATE "users"
		SET "firstName" = $1, "lastName" = $2, "updatedAt" = NOW()
		WHERE "id" = $3
	`
	_, err := dbClient.Exec(query, firstName, lastName, userID)
	if err != nil {
		return err
	}
	return nil
}
