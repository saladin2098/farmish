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
	Birth      string
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

type Schedule struct {
	ID    int
	Time1 time.Time
	Time2 time.Time
	Time3 time.Time
}

type ScheduleCreate struct {
	Time1 string
	Time2 string
	Time3 string
}

type FeedingSchedule struct {
	ID           int
	AnimalType   string `json:"animal_type"`
	LastFedIndex int  `json:"last_fed_index"`
	NextFedIndex int  `json:"next_fed_index"`
	ScheduleID   int `json:"schedule_id"`
}

type FeedingScheduleCreate struct {
	AnimalType   string
	LastFedIndex string
	NextFedIndex string
	ScheduleID   int
}
