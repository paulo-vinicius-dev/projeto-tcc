package service

import (
	"tcc-project/src/model"
	"tcc-project/src/repository"
)

// ReadAllUsers ...
func ReadAllUsers() ([]model.User, error) {
	users, err := repository.ReadAllUsers()
	if err != nil {
		return users, err
	}
	return users, nil
}

// ReadUserByID ...
func ReadUserByID(ID int) (model.User, error) {
	user, err := repository.ReadUserByID(ID)
	if err != nil {
		return user, err
	}
	return user, nil
}

// CreateUser ...
func CreateUser(user model.User) error {
	if err := repository.VerifyDuplicities(user); err != nil {
		return err
	}
	if err := repository.CreateUser(user); err != nil {
		return err
	}
	return nil
}

// UpdateUser ...
func UpdateUser(user model.User) error {
	if err := repository.VerifyDuplicities(user); err != nil {
		return err
	}
	if err := repository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

// DeleteUser ...
func DeleteUser(ID int) error {
	if err := repository.DeleteUser(ID); err != nil {
		if err := repository.DeleteLogicUser(ID); err != nil {
			return err
		}
	}
	return nil
}
