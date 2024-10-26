package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"tcc-project/src/services"

	"github.com/gorilla/mux"
)

// ReadAllUnits ...
func ReadAllUnits(w http.ResponseWriter, r *http.Request) {
	units, err := services.ReadAllUnits()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(units)
}

// ReadUnitByID ...
func ReadUnitByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	unit, err := services.ReadUnitByID(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(unit)
}
