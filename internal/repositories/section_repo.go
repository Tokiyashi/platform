package repositories

import (
	"context"
	"fmt"
	"platform/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Section_repo_interface interface {
	GetSections() ([]models.Section, error)
	AddSection(section *models.Section) error
	DeleteOneSection(id int) error
	UpdateSection(section *models.Section) error
}

type Section_repo struct {
	pool *pgxpool.Pool
}

func NewSectionRepo(pool *pgxpool.Pool) *Section_repo {
	return &Section_repo{pool: pool}
}

func (r *Section_repo) GetSections(courseId string) ([]models.Section, error) {
	rows, err := r.pool.Query(context.Background(), "SELECT id, title, description, course_id FROM sections WHERE course_id = $1", courseId)
	if err != nil {
		return nil, fmt.Errorf("error fetching sections: %w", err)
	}
	defer rows.Close()

	var sections []models.Section
	for rows.Next() {
		var section models.Section
		err = rows.Scan(&section.Id, &section.Title, &section.Description, &section.CourseId)
		if err != nil {
			return nil, fmt.Errorf("error scanning section: %w", err)
		}
		sections = append(sections, section)
	}
	return sections, nil
}

func (r *Section_repo) AddSection(section *models.Section) error {
	_, err := r.pool.Exec(context.Background(),
		"INSERT INTO sections (title, description, course_id) VALUES ($1, $2, $3)",
		section.Title, section.Description, section.CourseId)
	if err != nil {
		return fmt.Errorf("error adding section: %w", err)
	}
	return nil
}

func (r *Section_repo) DeleteOneSection(id int) error {
	_, err := r.pool.Exec(context.Background(), "DELETE FROM sections WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting section: %w", err)
	}
	return nil
}

func (r *Section_repo) UpdateSection(section *models.Section) error {
	_, err := r.pool.Exec(context.Background(),
		"UPDATE sections SET title = $1, description = $2, course_id = $3 WHERE id = $4",
		section.Title, section.Description, section.CourseId, section.Id)
	if err != nil {
		return fmt.Errorf("error updating section: %w", err)
	}
	return nil
}
