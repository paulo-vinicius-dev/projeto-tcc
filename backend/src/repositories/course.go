package repositories

import (
	"context"
	"tcc-project/src/database"
	"tcc-project/src/models"
)

// ReadAllCourses ...
func ReadAllCourses() ([]models.Course, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var courses []models.Course

	rows, err := conn.Query(context.Background(),
		`SELECT name, description, owner FROM course WHERE deleted_at IS NOT NULL`,
	)
	if err != nil {
		return courses, err

	}
	defer rows.Close()

	for rows.Next() {
		var course models.Course
		if err := rows.Scan(&course.Name, &course.Description, &course.Owner); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
