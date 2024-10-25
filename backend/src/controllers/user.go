package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tcc-project/src/models"
	"tcc-project/src/services"

	"github.com/gorilla/mux"
)

// TODO: This router ReadAllUser has bugs they need to be fix

// ReadAllUsers ...
func ReadAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.ReadAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, err)
		return
	}

	// TODO: find a place to this, I think it's out of place here
	JSONUsers, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(JSONUsers)
}

// ReadUserByID ...
func ReadUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, err := services.ReadUserByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, err)
		return
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
	}

	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJSON)
}

// CreateUser ...
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	if err := user.ValidateFields(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	if err := services.CreateUser(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// UpdateUser ...
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user.ID = ID

	if err := user.ValidateFields(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if err := services.UpdateUser(user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

// DeleteUser ...
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	if err := services.DeleteUser(ID); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
