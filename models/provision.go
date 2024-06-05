package models

import "time"

type Provision struct {
	ID         int
	Type       string
	AnimalType string
	Quantity   float64
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  int
}

type CreateProvision struct {
	ID         int
	Type       string
	AnimalType string
	Quantity   float64
}

type BodyProvision struct {
	Type       string
	AnimalType string
	Quantity   float64
}

type UpdateProvision struct {
	ID       int
	Type     string
	Quantity float64
}

type Filter struct {
	Limit  int
	OFFSET int
}
type GetProvision struct {
	ID       int
	Type     string
	Quantity float64
}

type GetAllProvisions struct {
	Count      int16
	Provisions *[]GetProvision
}
