package service

import (
	"farmish/config"
	"farmish/models"
	"farmish/postgres/managers"
)

func (p *Service) CreateProvision(prs *models.BodyProvision) (*models.CreateProvision, error) {
	provisionRepo := managers.NewProvisionRepo(p.db)

	ids, err := provisionRepo.GetAllIds()
	if err != nil {
		return nil, err
	}

	id, err := config.GenNewID(ids)
	if err != nil {
		return nil, err
	}
	createdProvision := &models.CreateProvision{
		ID:         id,
		Type:       prs.Type,
		AnimalType: prs.AnimalType,
		Quantity:   prs.Quantity,
	}

	createdProvision, err = provisionRepo.CreateProvision(*createdProvision)
	if err != nil {
		return nil, err
	}

	return createdProvision, nil
}

func (p *Service) GetProvision(id int, typ, animal_type string, quantity float64) (*models.GetProvision, error) {
	provisionRepo := managers.NewProvisionRepo(p.db)

	provision, err := provisionRepo.GetProvision(id, typ, animal_type, quantity)

	return provision, err
}

func (p *Service) GetAllProvision(filter models.Filter) (*models.GetAllProvisions, error) {
	provisionRepo := managers.NewProvisionRepo(p.db)

	provisions, err := provisionRepo.GetAllProvision(filter)

	return provisions, err
}

func (p *Service) UpdateProvision(prs *models.UpdateProvision) (*models.UpdateProvision, error) {
	provisionRepo := managers.NewProvisionRepo(p.db)

	oldPrs, err := provisionRepo.GetProvision(prs.ID, "", "", 0)
	if err != nil {
		return nil, err
	}

	var updatePrs = models.UpdateProvision{
		ID: oldPrs.ID,
	}

	if prs.Type == "" {
		updatePrs.Type = oldPrs.Type
	} else {
		updatePrs.Type = prs.Type
	}
	if prs.Quantity != 0 {
		updatePrs.Quantity = oldPrs.Quantity
	} else {
		updatePrs.Quantity = prs.Quantity
	}

	err = provisionRepo.UpdateProvision(&updatePrs)

	return &updatePrs, err
}

func (p *Service) DeleteProvision(id int) error {
	provisionRepo := managers.NewProvisionRepo(p.db)

	err := provisionRepo.DeleteProvision(id)

	return err
}
