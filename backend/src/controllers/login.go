package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"tcc-project/src/models"
	"tcc-project/src/services"
)

// Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	var login models.LoginDTO

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := login.ValidateFields(); err != nil {
		http.Error(w, fmt.Sprintf("Validation error: %v", err), http.StatusBadRequest)
		return
	}

	token, ok, err := services.Login(login)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error during login: %v", err), http.StatusInternalServerError)
		return
	}

	if !ok {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Logout ...
func Logout(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	services.Logout(token)

	w.WriteHeader(http.StatusOK)
}
