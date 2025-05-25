package repositories

import (
	"context"
	"fmt"
	"platform/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseRepository interface {
	GetAllCourses(ctx context.Context) ([]models.Course, error)
	GetCourseByID(ctx context.Context, id int) (models.Course, error)
	GetUserCourses(userID string) ([]*models.Course, error)
	AddCourse(ctx context.Context, course models.Course) (int, error)
	UpdateCourse(ctx context.Context, course models.Course) error
	DeleteCourse(ctx context.Context, id int) error
	JoinCourse(ctx context.Context, userID string, courseID int) error
}

type CourseRepo struct {
	pool *pgxpool.Pool
}

func NewCourseRepo(pool *pgxpool.Pool) *CourseRepo {
	return &CourseRepo{pool: pool}
}

func (r *CourseRepo) GetAllCourses(ctx context.Context) ([]models.Course, error) {
	rows, err := r.pool.Query(ctx, "SELECT * FROM courses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []models.Course
	for rows.Next() {
		var course models.Course
		err = rows.Scan(&course.Id, &course.Title, &course.Description, &course.Creator_id, &course.Created_at)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (r *CourseRepo) GetCourseByID(ctx context.Context, id int) (models.Course, error) {
	var course models.Course
	err := r.pool.QueryRow(ctx, "SELECT id, title, description FROM courses WHERE id = $1", id).
		Scan(&course.Id, &course.Title, &course.Description)
	if err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (r *CourseRepo) AddCourse(ctx context.Context, course models.Course) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		"INSERT INTO courses (title, description, creator_id) VALUES ($1, $2, $3) RETURNING id",
		course.Title, course.Description, course.Creator_id).Scan(&id)
	if err != nil {
		fmt.Print(err)
		return 0, err
	}
	return id, nil
}

func (r *CourseRepo) JoinCourse(ctx context.Context, userID int, courseID int) error {
	_, err := r.pool.Exec(ctx, "INSERT INTO users_courses (user_id, course_id) VALUES ($1, $2)", userID, courseID)
	return err
}

func (r *CourseRepo) UpdateCourse(ctx context.Context, course models.Course) error {
	_, err := r.pool.Exec(ctx,
		"UPDATE courses SET title = $1, description = $2 WHERE id = $3",
		course.Title, course.Description, course.Id)
	return err
}

func (r *CourseRepo) DeleteCourse(ctx context.Context, id int) error {
	_, err := r.pool.Exec(ctx, "DELETE FROM courses WHERE id = $1", id)
	return err
}

func (db *CourseRepo) GetUserCourses(userID string) ([]*models.Course, error) {
	query := `
        SELECT c.id, c.title, c.description
        FROM courses c
        INNER JOIN users_courses uc ON c.id = uc.course_id
        WHERE uc.user_id = $1
    `

	rows, err := db.pool.Query(context.Background(), query, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching user courses: %w", err)
	}
	defer rows.Close()

	var courses []*models.Course
	for rows.Next() {
		course := &models.Course{}
		err := rows.Scan(&course.Id, &course.Title, &course.Description)
		if err != nil {
			return nil, fmt.Errorf("error scanning course: %w", err)
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating courses: %w", err)
	}

	return courses, nil
}
