package repository

import (
	"context"
	"errors"
	"tcc-project/src/database"
	"tcc-project/src/model"
)

// ReadAllUsers ...
func ReadAllUsers() ([]model.User, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var users []model.User

	rows, err := conn.Query(context.Background(),
		`SELECT id, name, email FROM users WHERE deleted_at IS NULL`,
	)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return users, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

// ReadUserByID ...
func ReadUserByID(ID int) (model.User, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var user model.User

	row := conn.QueryRow(context.Background(),
		`SELECT id, name, email FROM users WHERE deleted_at IS NULL AND id = $1`,
		ID,
	)

	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser ...
func CreateUser(user model.User) error {
	conn := database.Connection()
	defer conn.Close(context.Background())

	if _, err := conn.Exec(context.Background(),
		`INSERT INTO users(name, email, password) VALUES ($1, $2, $3)`,
		user.Name, user.Email, user.Password,
	); err != nil {
		return err
	}

	return nil
}

// UpdateUser ...
func UpdateUser(user model.User) error {
	conn := database.Connection()
	defer conn.Close(context.Background())
	if _, err := conn.Exec(context.Background(),
		`UPDATE users SET name = $1, email = $2 WHERE deleted_at IS NULL AND id = $3`,
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
		`DELETE FROM users WHERE deleted_at IS NULL AND id = $1`,
		ID,
	); err != nil {
		return err
	}
	return nil

}

// DeleteLogicUser ...
func DeleteLogicUser(ID int) error {
	conn := database.Connection()
	defer conn.Close(context.Background())
	if _, err := conn.Exec(context.Background(),
		`UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE deleted_at IS NULL AND id = $1`,
		ID,
	); err != nil {
		return err
	}
	return nil
}

// VerifyDuplicities ...
func VerifyDuplicities(user model.User) error {
	conn := database.Connection()
	defer conn.Close(context.Background())

	row := conn.QueryRow(context.Background(),
		`SELECT COUNT(id) FROM users WHERE email = $1 AND id != $2 AND deleted_at IS NULL`,
		user.Email, user.ID,
	)

	var count int

	if err := row.Scan(&count); err != nil || count > 0 {
		return errors.New("some user with same email already exists in database")
	}

	return nil
}
