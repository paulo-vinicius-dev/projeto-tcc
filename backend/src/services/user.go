package services

import (
	"tcc-project/src/models"
	"tcc-project/src/repositories"
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
	if err := repositories.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

// DeleteUser ...
func DeleteUser(ID int) error {
	if err := repositories.DeleteUser(ID); err != nil {
		if err := repositories.DeleteLogicUser(ID); err != nil {
			return err
		}
	}
	return nil
}
