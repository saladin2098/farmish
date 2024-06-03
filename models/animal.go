package models

import "time"

type Animal struct {
	ID              int32
	Type            string
	Birth           time.Time
	Weight          int32
	HealthCondition HealthCondition
	Feeding         FeedingSchedule
	AvgConsumption  float32
	AvgWater        float32

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt int
}

type AnimalGet struct {
	ID              int32
	Type            string
	Birth           time.Time
	Weight          int32
	HealthCondition HealthCondition
	Feeding         FeedingSchedule
	AvgConsumption  float32
	AvgWater        float32
}

type AnimalsGetAll struct {
	Count   int32
	Animals *[]AnimalGet
}

type AnimalCreate struct {
	ID         int32
	Type       string
	Birth      time.Time
	Weight     int32
	AnimalType string
	IsHealthy  bool
	Condition  string
	Medication string
}

type AnimalUpdate struct {
	ID         int32
	Weight     int32
	IsHealthy  bool
	Condition  string
	Medication string
}

type FeedingSchedule struct {
	ID           int
	LastFedINdex int
}