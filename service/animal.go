package service

import (
	"farmish/config"
	"farmish/postgres/managers"
	"time"
)

type AnimalService struct {
	HR *managers.HealthConditionRepo
	AR *managers.AnimalRepo
}

func (s *AnimalService) CreateAnimal() (error, error) {
	animalIDs, err := s.AR.GetAllAnimalIds()
	if err != nil {
		return err, nil
	}
	newID1, err := config.GenNewID(animalIDs)
	if err != nil {
		return err, nil
	}
	healthConditionIds, err := s.HR.GetAllHealthConditionIds()
	if err != nil {
		return err, nil
	}
	newID2, err := config.GenNewID(healthConditionIds)
	if err != nil {
		return err, nil
	}
	animal.ID = int32(newID1)
	newAnimal, err := s.AR.CreateAnimal(animal, float64(1), float64(2), newID2)
	if err != nil {
		return err, nil
	}
	birth, err := time.Parse(time.RFC3339, animal.Birth)
	if err != nil {
		return err, nil
	}
	duration := time.Since(birth)
	lived_days := int32(duration.Hours() / 24)
	consumptionPerDay := config.CalcWaterForPoultryPerDay(float64(lived_days))
	s.AR.UpdateAvgConsumption(int(newAnimal.ID), consumptionPerDay/3, float64(4))
	return nil, nil
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
