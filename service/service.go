package service

import (
	"database/sql"
	"farmish/postgres/managers"
)

type Service struct {
	db *sql.DB
	AS managers.AnimalRepo
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}
