package services

import (
	"tcc-project/src/models"
	"tcc-project/src/repositories"
)

// ReadAllCourses ...
func ReadAllCourses() ([]models.CourseDTO, error) {
	courses, err := repositories.ReadAllCourses()
	if err != nil {
		return courses, err
	}

	return courses, nil
}

// ReadCourseByID ...
func ReadCourseByID(ID int) (models.CourseDTO, error) {
	course, err := repositories.ReadCourseByID(ID)
	if err != nil {
		return course, err
	}
	return course, nil
}
