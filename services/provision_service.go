package service

import (
	"farmish/config"
	"farmish/models"
	"farmish/postgresql/managers"
	"fmt"
)

type ProvisionService struct {
	PR *managers.ProvisionRepo
}

func NewProvisionService(pr *managers.ProvisionRepo) *ProvisionService {
	return &ProvisionService{PR: pr}
}

func (p *ProvisionService) CreateProvision(prs *models.BodyProvision) error {

	ids, err := p.PR.GetAllIds()
	fmt.Println(ids)
	if err != nil {
		return err
	}

	id, err := config.GenNewID(ids)
	if err != nil {
		return err
	}

	err = p.PR.CreateProvision(*prs, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProvisionService) GetProvisionById(id int) (*models.GetProvision, error) {
	provision, err := p.PR.GetProvisionById(id)

	return provision, err
}

func (p *ProvisionService) GetAllProvision(filter models.Filter, typ, animalTaye string, quantity float64) (*[]models.GetProvision, error) {
	provisions, err := p.PR.GetAllProvision(filter, typ, animalTaye, quantity)

	return &provisions, err
}

func (p *ProvisionService) UpdateProvision(prs *models.UpdateProvision, id int) (*models.UpdateProvision, error) {
	oldPrs, err := p.PR.GetProvisionById(id)
	if err != nil {
		return nil, err
	}

	var updatePrs = models.UpdateProvision{}

	if prs.Type == "" {
		updatePrs.Type = oldPrs.Type
	} else {
		updatePrs.Type = prs.Type
	}
	if prs.Quantity == 0 {
		updatePrs.Quantity = oldPrs.Quantity
	} else {
		updatePrs.Quantity = prs.Quantity
	}

	err = p.PR.UpdateProvision(&updatePrs, id)

	return &updatePrs, err
}

func (p *ProvisionService) DeleteProvision(id int) error {
	err := p.PR.DeleteProvision(id)

	return err
}

func (p *ProvisionService) ProvisionData(animal_type string) (bool, error) {
	b, err := p.PR.ProvisionData(animal_type)

	return b, err
}
