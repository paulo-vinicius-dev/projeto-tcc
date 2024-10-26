package services

import (
	"tcc-project/src/models"
	"tcc-project/src/repositories"
)

// ReadAllUnits ...
func ReadAllUnits() ([]models.UnitDTO, error) {
	units, err := repositories.ReadAllUnits()
	if err != nil {
		return units, err
	}
	return units, nil
}

// ReadUnitByID ...
func ReadUnitByID(ID int) (models.UnitDTO, error) {
	unit, err := repositories.ReadUnitByID(ID)
	if err != nil {
		return unit, err
	}
	return unit, nil
}
