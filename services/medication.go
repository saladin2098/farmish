package service

import (
	"farmish/config"
	"farmish/models"
	"farmish/postgresql"
)

type MedicationService struct {
	MedRepo *postgresql.MedicationRepo
}

func NewMedicationService(mr *postgresql.MedicationRepo) *MedicationService {
	return &MedicationService{MedRepo: mr}
}

func (s *MedicationService) CreateMedication(medication *models.MedicationsGet) (*models.Medications, error) {
	medIds, err := s.MedRepo.GetAllMedicationIDs()
	if err != nil {
		return nil, err
	}
	newID, err := config.GenNewID(*medIds)
	if err!= nil {
        return nil, err
    }
	var redCreate models.Medications 
	redCreate.ID = newID
	redCreate.Name = medication.Name
	redCreate.Type = medication.Type
	redCreate.Quantity = medication.Quantity
	return s.MedRepo.CreateMedication(&redCreate)
}

func (s *MedicationService) GetMedication(id int, name, turi string) (*models.Medications, error) {
    return s.MedRepo.GetMedication(id, name, turi)
}

func (s *MedicationService) DeleteMedication(id int) error {
    return s.MedRepo.DeleteMedication(id)
}

func (s *MedicationService) UpdateMedication(med *models.Medications) error {
    return s.MedRepo.UpdateMedication(med)
}

func (s *MedicationService) GetMedicationsGroupedByType(tur string) (*models.MedicarionaGrouped, error) {
    return s.MedRepo.GetMedicationsGroupedByType(tur)
}