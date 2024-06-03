package service

import (
	"database/sql"
	"farmish/models"
	"farmish/postgres/managers"
)

type MedicationService struct {
	medicationRepo *managers.MedicationRepo
}

func NewMedicationService(db *sql.DB) *MedicationService {
	return &MedicationService{managers.NewMidacationRepo(db)}
}

func (s *MedicationService) CreateMedication(medication *models.Medications) (*models.Medications, error) {
	return nil, nil
}
