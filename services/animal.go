package service

import (
	"farmish/config"
	"farmish/models"
	"farmish/postgresql/managers"
	"math"
	"time"
)

type AnimalService struct {
	HR *managers.HealthConditionRepo
	AR *managers.AnimalRepo
}

func NewAnimalService(hr *managers.HealthConditionRepo, ar *managers.AnimalRepo) *AnimalService {
	return &AnimalService{HR: hr, AR: ar}
}

func (s *AnimalService) CreateAnimal(animal *models.AnimalCreate) error {
	animalIDs, err := s.AR.GetAllAnimalIds()
	if err != nil {
		return err
	}
	newID1, err := config.GenNewID(animalIDs)
	if err != nil {
		return err
	}
	healthConditionIds, err := s.HR.GetAllHealthConditionIds()
	if err != nil {
		return err
	}
	newID2, err := config.GenNewID(healthConditionIds)
	if err != nil {
		return err
	}
	animal.ID = int32(newID1)
	newAnimal, err := s.AR.CreateAnimal(animal, float64(1), float64(2), newID2)
	if err != nil {
		return err
	}
	birth, err := time.Parse("2006-01-02", animal.Birth)
	if err != nil {
		return err
	}
	duration := time.Since(birth)
	lived_days := int32(duration.Hours() / 24)

	// lived_days, err := s.AR.GetAnimalAgeInDays(int(newAnimal.ID))
	// if err != nil {
	// 	return err
	// }

	if newAnimal.AnimalType == "parranda" {
		wc, fc := config.CalcConsumptionForPoultry(float64(lived_days))
		err := s.AR.UpdateAvgConsumption(int(newAnimal.ID), math.Round((wc/3)*1000)/1000, math.Round((fc/3)*1000)/1000)
		if err != nil {
			return err
		}
	} else if newAnimal.AnimalType == "hayvon" {
		wc, fc := config.CalcConsumptionForAnimals(float64(lived_days), float64(newAnimal.Weight))
		err := s.AR.UpdateAvgConsumption(int(newAnimal.ID), math.Round((wc/3)*1000)/1000, math.Round((fc/3)*1000)/1000)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *AnimalService) GetAnimal(id int) (*models.Animal, error) {
	animal, err := s.AR.GetAnimalByID(id)
	if err != nil {
		return nil, err
	}
	return animal, nil
}

func (s *AnimalService) UpdateAnimal(animal *models.AnimalUpdate) error {
	if animal.IsHealthy {
		animal.Condition = "Healthy"
		animal.Medication = "None"
	}
	err := s.AR.UpdateAnimal(animal)
	if err != nil {
		return err
	}
	return nil
}

func (s *AnimalService) DeleteAnimal(id int) error {
	err := s.AR.DeleteAnimal(id)
	if err != nil {
		return err
	}
	return nil
}

/*
	### Average Weight of a 1-Year-Old Sheep
	- Sheep: Typically, a one-year-old sheep (also referred to as a yearling) weighs between 25 to 45 kg (55 to 100 lbs),
			depending on the breed and whether it’s a male (ram) or female (ewe).

	### Average Weight of a 1-Year-Old Cow
	- Cow: A one-year-old cow (technically called a yearling calf) can weigh quite a bit more, generally ranging from 200 to 350 kg
		(440 to 770 lbs), again depending on the breed and sex.

*/

/*
	### Average Weight of a 2-Month-Old Sheep
	- Sheep: A lamb (baby sheep) at 2 months old typically weighs between 12 to 18 kg (26 to 40 lbs),
		depending on the breed and whether it’s a male or female.

	### Average Weight of a 2-Month-Old Cow
	- Cow: A calf (young cow) at 2 months old generally weighs between 70 to 90 kg (154 to 198 lbs),
		depending on the breed, sex, and overall health.

*/
