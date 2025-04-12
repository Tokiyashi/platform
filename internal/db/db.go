package db

import (
	"context"
	"log"
	"platform/internal/repositories"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	pool     *pgxpool.Pool
	Users    repositories.UserRepo
	Courses  repositories.CourseRepo
	Sections repositories.Section_repo
}

func New(connstring string) (*DB, error) {
	pool, err := pgxpool.New(context.Background(), connstring)
	if err != nil {
		log.Println(("Error conn to db"))
		return nil, err
	}
	err = pool.Ping(context.Background())
	if err != nil {
		log.Println(("Error conn to db"))
		return nil, err
	}
	ur := *repositories.NewUserRepo(pool)
	cr := *repositories.NewCourseRepo(pool)

	return &DB{pool: pool,
		Users:   ur,
		Courses: cr,
	}, nil
}
