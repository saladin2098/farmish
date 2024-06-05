package dashboard

import (
	"farmish/models"
	service "farmish/services"
)

type Dashboard struct {
	Service service.Service
}

func NewDashboard(service service.Service) *Dashboard {
	return &Dashboard{Service: service}
}

// animals int, poultries int
func (d *Dashboard) GetAnimalsCount() (int, int, error) {
	animals, err := d.Service.AR.GetAllAnimals("hayvon", "", "")
	if err != nil {
		panic(err)
	}
	poultries, err := d.Service.AR.GetAllAnimals("parranda", "", "")
	if err != nil {
		panic(err)
	}
	return int(animals.Count), int(poultries.Count), nil
}

// animals int, poultries int
func (d *Dashboard) GetAvgWeight() (float32, float32, error) {
	animals, err := d.Service.AR.GetAllAnimals("hayvon", "", "")
	if err != nil {
		return 0, 0, err
	}
	poultries, err := d.Service.AR.GetAllAnimals("parranda", "", "")
	if err != nil {
		return 0, 0, err
	}

	var animal_avg, poultrie_avg float32
	for _, v := range animals.Animals {
		animal_avg += float32(v.Weight)
	}
	for _, v := range poultries.Animals {
		poultrie_avg += float32(v.Weight)
	}
	return animal_avg / float32(animals.Count), poultrie_avg / float32(poultries.Count), nil
}

func (d *Dashboard) GetSickAnimals() (*models.AnimalsGetAll, error) {
	result := models.AnimalsGetAll{}
	animals, err := d.Service.AR.GetAllAnimals("", "false", "")
	if err != nil {
		return nil, err
	}
	result.Animals = append(result.Animals, animals.Animals...)
	result.Count = animals.Count
	return &result, nil
}

func (d *Dashboard) GetHungryAnimals() (*models.AnimalsGetAll, error) {
	result := models.AnimalsGetAll{}
	animals, err := d.Service.AR.GetAllAnimals("", "", "true")
	if err != nil {
		return nil, err
	}
	result.Animals = append(result.Animals, animals.Animals...)
	result.Count = animals.Count
	return &result, nil
}

func (d *Dashboard) CheckProvision() (bool, bool, error) {
	pr_anim, err := d.Service.PS.ProvisionData("hayvon")
	if err != nil {
		return false, false, err
	}
	pr_poul, err := d.Service.PS.ProvisionData("parranda")
	if err != nil {
		return false, false, err
	}

	return pr_anim, pr_poul, nil
}
