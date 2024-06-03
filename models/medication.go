package models

type Medications struct{
	ID int
	Name string
	Type string
	Quantity float32
}


type MedicinesGetAll struct{
	Count int32
	Medications *[]Medications
}

type HealthCondition struct {
	ID        int
	AnimalID  int
	IsHealthy bool
	Condition string
	Medication string
	IsTreated bool
 }