package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{DB: db}
}

func (c *Course) Create(name, description, categoryId string) (Course, error) {
	id := uuid.New().String()

	_, err := c.DB.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryId)
	if err != nil {
		return Course{}, err
	}

	return Course{
		ID:          id,
		Name:        name,
		Description: description,
		CategoryID:  categoryId,
	}, nil
}

func (c *Course) GetAll() ([]Course, error) {
	rows, err := c.DB.Query("SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID); err != nil {
			return nil, err
		}
		result = append(result, course)
	}

	return result, nil
}

func (c *Course) FindByCategoryId(categoryId string) ([]Course, error) {
	rows, err := c.DB.Query("SELECT * FROM courses WHERE category_id = $1", categoryId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID); err != nil {
			return nil, err
		}
		result = append(result, course)
	}

	return result, nil
}
