package server

import (
	"fmt"
	"log"
	"os"
	"platform/internal/api"
	"platform/internal/db"
)

type Server struct {
	db  *db.DB
	api *api.API
}

// const connstring = "postgres://postgres:postgres@db:5432/platform?sslmode=disable"

func New() (*Server, error) {
	var connstring string = os.Getenv("DATABASE_URL")

	s := &Server{}
	log.Print(connstring)
	db, err := db.New(connstring)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	s.db = db
	s.api = api.New(db)

	return s, nil
}

func (s *Server) Run() {
	s.api.Start()
}
