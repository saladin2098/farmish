package service

import (
	"farmish/config"
	"farmish/models"
	"farmish/postgres/managers"
)

type ProvisionService struct {
	PR *managers.ProvisionRepo
}

func (p *ProvisionService) CreateProvision(prs *models.BodyProvision) (*models.CreateProvision, error) {

	ids, err := p.PR.GetAllIds()
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

	createdProvision, err = p.PR.CreateProvision(*createdProvision)
	if err != nil {
		return nil, err
	}

	return createdProvision, nil
}

func (p *ProvisionService) GetProvision(id int, typ, animal_type string, quantity float64) (*models.GetProvision, error) {
	provision, err := p.PR.GetProvision(id, typ, animal_type, quantity)

	return provision, err
}

func (p *ProvisionService) GetAllProvision(filter models.Filter) (*models.GetAllProvisions, error) {
	provisions, err := p.PR.GetAllProvision(filter)

	return provisions, err
}

func (p *ProvisionService) UpdateProvision(prs *models.UpdateProvision) (*models.UpdateProvision, error) {
	oldPrs, err := p.PR.GetProvision(prs.ID, "", "", 0)
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

	err = p.PR.UpdateProvision(&updatePrs)

	return &updatePrs, err
}

func (p *ProvisionService) DeleteProvision(id int) error {
	err := p.PR.DeleteProvision(id)

	return err
}
