package models

type Medications struct{
	ID int
	Name string
	Type string
	Quantity float32
}

type MedicationsGet struct{
	Name string
	Type string
	Quantity float32
}


type MedicinesGetAll struct{
	Type string
	Count int
	Medications []MedicationsGet
}
type MedicarionaGrouped struct {
	Medications []MedicinesGetAll
}


type HealthCondition struct {
	ID        int
	AnimalID  int
	IsHealthy bool
	Condition string
	Medication string
	IsTreated bool
}