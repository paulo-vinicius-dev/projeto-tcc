package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tcc-project/src/services"

	"github.com/gorilla/mux"
)

// ReadAllCourses ...
func ReadAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := services.ReadAllCourses()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)
}

// ReadCourseByID ...
func ReadCourseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	course, err := services.ReadCourseByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("Content-type", "aplication/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(course)
}
