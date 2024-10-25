package services

import (
	"tcc-project/src/models"
	"tcc-project/src/repositories"
)

// ReadAllCourses
func ReadAllCourses() ([]models.Course, error) {
	courses, err := repositories.ReadAllCourses()
	if err != nil {
		return courses, err
	}

	return courses, nil
}
