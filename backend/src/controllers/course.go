package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tcc-project/src/services"
)

// ReadAllCourses ...
func ReadAllCourses(w http.ResponseWriter, r *http.Request) {
	courses, err := services.ReadAllCourses()
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprint(w, err)
		return
	}

	JSONCourses, err := json.Marshal(&courses)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/Json")
	w.WriteHeader(http.StatusOK)
	w.Write(JSONCourses)
}
