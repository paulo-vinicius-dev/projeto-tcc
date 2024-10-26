package repositories

import (
	"context"
	"tcc-project/src/database"
	"tcc-project/src/models"
)

// ReadAllUnits ...
func ReadAllUnits() ([]models.UnitDTO, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var units []models.UnitDTO

	rows, err := conn.Query(context.Background(),
		`SELECT 
			u.id AS unit_id,
			u.level AS unit_level,
			u.order AS unit_order,
			u.name AS unit_name,
			u.description AS unit_description,
			c.id AS course_id,
			c.name AS course_name,
			c.description AS course_description,
			c.access_type AS course_access_type,
			c.access_code AS course_access_code,
			u2.id AS user_id,
			u2.name AS user_name,
			u2.email AS user_email	
		FROM unit u
		INNER JOIN course c ON u.course = c.id
		INNER JOIN "user" u2 ON c.owner = u2.id   
		WHERE u.deleted_at IS NULL`,
	)
	if err != nil {
		return units, err
	}
	defer rows.Close()

	for rows.Next() {
		var unit models.UnitDTO
		if err := rows.Scan(
			&unit.ID,
			&unit.Level,
			&unit.Order,
			&unit.Name,
			&unit.Description,
			&unit.Course.ID,
			&unit.Course.Name,
			&unit.Course.Description,
			&unit.Course.AccessType,
			&unit.Course.AccessCode,
			&unit.Course.Owner.ID,
			&unit.Course.Owner.Name,
			&unit.Course.Owner.Email,
		); err != nil {
			return units, err
		}
		units = append(units, unit)
	}

	return units, nil
}

// ReadUnitByID ...
func ReadUnitByID(ID int) (models.UnitDTO, error) {
	conn := database.Connection()
	defer conn.Close(context.Background())

	var unit models.UnitDTO

	row := conn.QueryRow(context.Background(),
		`SELECT 
			u.id AS unit_id,
			u.level AS unit_level,
			u.order AS unit_order,
			u.name AS unit_name,
			u.description AS unit_description,
			c.id AS course_id,
			c.name AS course_name,
			c.description AS course_description,
			c.access_type AS course_access_type,
			c.access_code AS course_access_code,
			u2.id AS user_id,
			u2.name AS user_name,
			u2.email AS user_email	
		FROM unit u
		INNER JOIN course c ON u.course = c.id
		INNER JOIN "user" u2 ON c.owner = u2.id   
		WHERE u.deleted_at IS NULL AND u.id = $1`,
		ID,
	)
	if err := row.Scan(
		&unit.ID,
		&unit.Level,
		&unit.Order,
		&unit.Name,
		&unit.Description,
		&unit.Course.ID,
		&unit.Course.Name,
		&unit.Course.Description,
		&unit.Course.AccessType,
		&unit.Course.AccessCode,
		&unit.Course.Owner.ID,
		&unit.Course.Owner.Name,
		&unit.Course.Owner.Email,
	); err != nil {
		return unit, err
	}
	return unit, nil
}
