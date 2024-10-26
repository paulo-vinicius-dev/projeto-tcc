package repositories

import (
	"context"
	"errors"
	"tcc-project/src/database"
	"tcc-project/src/models"
)

// ReadAllCourses ...
func ReadAllCourses() ([]models.CourseDTO, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var courses []models.CourseDTO

	rows, err := conn.Query(context.Background(),
		`SELECT 
			c.id AS course_id,
			c.name AS course_name,
			c.description AS course_description,
			c.access_type AS course_access_type,
			c.access_code AS course_access_code,
			u2.id AS user_id,
			u2.name AS user_name,
			u2.email AS user_email	
		FROM course c 
		INNER JOIN "user" u2 ON c.owner = u2.id   
		WHERE c.deleted_at IS NULL`,
	)
	if err != nil {
		return courses, err

	}
	defer rows.Close()

	for rows.Next() {
		var course models.CourseDTO
		if err := rows.Scan(
			&course.ID,
			&course.Name,
			&course.Description,
			&course.AccessType,
			&course.AccessCode,
			&course.Owner.ID,
			&course.Owner.Name,
			&course.Owner.Email,
		); err != nil {
			return courses, err
		}
		courses = append(courses, course)
	}

	if len(courses) < 1 {
		return courses, errors.New("any course was found")
	}

	return courses, nil
}

// ReadCourseByID ...
func ReadCourseByID(ID int) (models.CourseDTO, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var course models.CourseDTO

	row := conn.QueryRow(context.Background(),
		`SELECT 
			c.id AS course_id,
			c.name AS course_name,
			c.description AS course_description,
			c.access_type AS course_access_type,
			c.access_code AS course_access_code,
			u2.id AS user_id,
			u2.name AS user_name,
			u2.email AS user_email	
		FROM course c 
		INNER JOIN "user" u2 ON c.owner = u2.id 
		WHERE c.deleted_at IS NULL AND c.id = $1`,
		ID,
	)

	if err := row.Scan(
		&course.ID,
		&course.Name,
		&course.Description,
		&course.AccessType,
		&course.AccessCode,
		&course.Owner.ID,
		&course.Owner.Name,
		&course.Owner.Email,
	); err != nil {
		return course, err
	}

	return course, nil
}
