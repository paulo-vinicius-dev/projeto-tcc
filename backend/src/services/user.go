package services

import (
	"tcc-project/src/models"
	"tcc-project/src/repositories"

	"golang.org/x/crypto/bcrypt"
)

// ReadAllUsers ...
func ReadAllUsers() ([]models.UserDTO, error) {
	users, err := repositories.ReadAllUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

// ReadUserByID ...
func ReadUserByID(ID int) (models.UserDTO, error) {
	user, err := repositories.ReadUserByID(ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser ...
func CreateUser(user models.User) error {
	if err := repositories.VerifyDuplicities(user); err != nil {
		return err
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := repositories.CreateUser(user); err != nil {
		return err
	}
	return nil
}

// UpdateUser ...
func UpdateUser(user models.User) error {
	if err := repositories.VerifyDuplicities(user); err != nil {
		return err
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := repositories.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

// DeleteUser ...
func DeleteUser(ID int) error {
	if err := repositories.DeleteUser(ID); err != nil {
		if err := repositories.SoftDeleteUser(ID); err != nil {
			return err
		}
	}
	return nil
}

// HashPassword ...
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CheckPassword ...
func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
