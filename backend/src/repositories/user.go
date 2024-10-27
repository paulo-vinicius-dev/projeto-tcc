package repositories

import (
	"context"
	"errors"
	"tcc-project/src/database"
	"tcc-project/src/models"
)

// ReadAllUsers ...
func ReadAllUsers() ([]models.UserDTO, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var users []models.UserDTO

	rows, err := conn.Query(context.Background(),
		`SELECT id, name, email FROM "user" WHERE deleted_at IS NULL`,
	)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserDTO
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if len(users) < 1 {
		return users, errors.New("any user was found")
	}

	return users, nil
}

// ReadUserByID ...
func ReadUserByID(ID int) (models.UserDTO, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var user models.UserDTO

	row := conn.QueryRow(context.Background(),
		`SELECT id, name, email FROM "user" WHERE deleted_at IS NULL AND id = $1`,
		ID,
	)

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser ...
func CreateUser(user models.User) error {
	conn := database.Connection()
	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(),
		`INSERT INTO "user"(name, email, password) VALUES ($1, $2, $3)`,
		user.Name, user.Email, user.Password,
	); err != nil {
		return err
	}

	return nil
}

// UpdateUser ...
func UpdateUser(user models.User) error {
	conn := database.Connection()
	defer conn.Close(context.Background())
	if _, err := conn.Exec(context.Background(),
		`UPDATE "user" SET name = $1, email = $2, updated_at = CURRENT_TIMESTAMP WHERE deleted_at IS NULL AND id = $3`,
		user.Name, user.Email, user.ID,
	); err != nil {
		return err
	}
	return nil
}

// DeleteUser ...
func DeleteUser(ID int) error {
	conn := database.Connection()
	defer conn.Close(context.Background())
	if _, err := conn.Exec(context.Background(),
		`DELETE FROM "user" WHERE deleted_at IS NULL AND id = $1`,
		ID,
	); err != nil {
		return err
	}
	return nil

}

// SoftDeleteUser ...
func SoftDeleteUser(ID int) error {
	conn := database.Connection()
	defer conn.Close(context.Background())
	if _, err := conn.Exec(context.Background(),
		`UPDATE "user" SET deleted_at = CURRENT_TIMESTAMP WHERE deleted_at IS NULL AND id = $1`,
		ID,
	); err != nil {
		return err
	}
	return nil
}

// VerifyDuplicities ...
func VerifyDuplicities(user models.User) error {
	conn := database.Connection()
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(),
		`SELECT COUNT(id) FROM "user" WHERE email = $1 AND id != $2 AND deleted_at IS NULL`,
		user.Email, user.ID,
	)

	var count int

	if err := row.Scan(&count); err != nil || count > 0 {
		return errors.New("some user with same email already exists in database")
	}

	return nil
}

// GetIDAndPasswordByEmail ...
func GetIDAndPasswordByEmail(email string) (int, string, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(),
		`SELECT id, password FROM "user" WHERE deleted_at IS NULL AND email = $1`,
		email,
	)
	var (
		ID       int
		password string
	)
	if err := row.Scan(&ID, &password); err != nil {
		return 0, "", err
	}
	return ID, password, nil
}
