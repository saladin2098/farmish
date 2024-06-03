package service

import (
	"database/sql"
	"farmish/models"
	"farmish/postgres/managers"
)

type AnimalService struct {
	animalRepo *managers.AnimalRepo
}

func NewAnimalService(db *sql.DB) *AnimalService {
	return &AnimalService{managers.NewAnimalRepo(db)}
}

func (s *AnimalService) CreateAnimal(animal *models.AnimalCreate) (*models.Animal, error) {

	return nil, nil
}
