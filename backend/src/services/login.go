package services

import (
	"tcc-project/src/auth"
	"tcc-project/src/middlewares"
	"tcc-project/src/models"
	"tcc-project/src/repositories"
)

// Login ...
func Login(login models.LoginDTO) (string, bool, error) {

	ID, password, err := repositories.GetIDAndPasswordByEmail(login.Email)
	if err != nil {
		return "", false, err
	}

	ok := CheckPassword(login.Password, password)
	if !ok {
		return "", false, nil
	}

	token, err := auth.GenerateJWT(ID)
	if err != nil {
		return "", false, err
	}

	return token, ok, nil
}

// Logout ...
func Logout(token string) {
	middlewares.BlacklistedToken[token] = struct{}{}
}
