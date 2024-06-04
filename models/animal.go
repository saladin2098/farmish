package models

import "time"

type Animal struct {
	ID              int32
	Type            string
	Birth           string
	Weight          int32
	AnimalType      string
	HealthCondition HealthCondition
	Feeding         FeedingSchedule
	AvgConsumption  float32
	AvgWater        float32

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt int
}

type AnimalGet struct {
	ID              int32  `json:"id"`
	Type            string `json:"type"`
	AnimalType      string `json:"animal_type"`
	Birth           string `json:"birth"`
	Weight          int32  `json:"weight"`
	HealthCondition HealthCondition
	Feeding         FeedingSchedule
	AvgConsumption  float32 `json:"avg_consumption"`
	AvgWater        float32 `json:"avg_water"`
}

type AnimalsGetAll struct {
	Count   int32
	Animals *[]AnimalGet
}

type AnimalCreate struct {
	ID         int32  `json:"id"`
	Type       string `json:"type"`
	Birth      string `json:"birth"`
	Weight     int32  `json:"weight"`
	AnimalType string `json:"animal_type"`
	IsHealthy  bool   `json:"is_healthy"`
	Condition  string `json:"condition"`
	Medication string `json:"medication"`
}

type AnimalUpdate struct {
	ID         int32  `json:"id"`
	Weight     int32  `json:"weight"`
	IsHealthy  bool   `json:"is_healthy"`
	Condition  string `json:"condition"`
	Medication string `json:"medication"`
}

type Schedule struct {
	ID         int
	AnimalType string
	Time1      time.Time
	Time2      time.Time
	Time3      time.Time
}

type FeedingSchedule struct {
	ID           int
	AnimalType   string
	NextFedIndex int
	LastFedIndex int
	ScheduleID int
}
