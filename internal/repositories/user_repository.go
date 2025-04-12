package repositories

import (
	"context"
	"fmt"
	"log"

	"platform/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepo struct {
	pool *pgxpool.Pool
}

func NewUserRepo(pool *pgxpool.Pool) *UserRepo {
	return &UserRepo{pool: pool}
}

func (db *UserRepo) GetUser(id string) (*models.User, error) {
	rows, err := db.pool.Query(context.Background(), "SELECT id, first_name, last_name, email FROM users WHERE id = $1", id)
	if err != nil {
		log.Println("user not found")
		return nil, err
	}
	defer rows.Close()

	user := &models.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (db *UserRepo) Auth(email, pass string) (*models.User, error) {
	rows, err := db.pool.Query(context.Background(), "SELECT id, first_name, last_name, email FROM users WHERE email = $1 AND password = $2", email, pass)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	user := &models.User{}
	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}

	err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (db *UserRepo) AddUser(firstName, lastName, email, pass string) error {
	_, err := db.pool.Exec(context.Background(), "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)", firstName, lastName, email, pass)
	if err != nil {
		return fmt.Errorf("ошибка при создании пользователя: %w", err)
	}
	return nil
}
