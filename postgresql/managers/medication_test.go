package managers

import (
	"farmish/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetMedication(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMidacationRepo(db)

	med := &models.Medications{ID: 1, Name: "Med1", Type: "Type1", Quantity: 10}
	rows := sqlmock.NewRows([]string{"id", "name", "type", "quantity"}).
		AddRow(med.ID, med.Name, med.Type, med.Quantity)
	mock.ExpectQuery(`select id, name, type, quantity from medications where id = \$1`).
		WithArgs(med.ID).
		WillReturnRows(rows)

	retrievedMed, err := repo.GetMedication(med.ID, "", "")
	assert.NoError(t, err)
	assert.Equal(t, med, retrievedMed)
}
func TestCreateMedication(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMidacationRepo(db)

	med := &models.Medications{ID: 1, Name: "Med1", Type: "Type1", Quantity: 10}
	rows := sqlmock.NewRows([]string{"id", "name", "type", "quantity"}).
		AddRow(med.ID, med.Name, med.Type, med.Quantity)
	mock.ExpectQuery(`insert into medications`).
		WithArgs(med.ID, med.Name, med.Type, med.Quantity).
		WillReturnRows(rows)

	createdMed, err := repo.CreateMedication(med)
	assert.NoError(t, err)
	assert.Equal(t, med, createdMed)
}

func TestDeleteMedication(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMidacationRepo(db)

	mock.ExpectExec(`delete from medications where id = \$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteMedication(1)
	assert.NoError(t, err)
}

func TestUpdateMedication(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMidacationRepo(db)

	med := &models.Medications{ID: 1, Name: "Med1", Type: "Type1", Quantity: 10}
	mock.ExpectExec(`update medications set name = \$1, type = \$2, quantity = \$3 where id = \$4`).
		WithArgs(med.Name, med.Type, med.Quantity, med.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateMedication(med)
	assert.NoError(t, err)
}
func TestGetMedicationsGroupedByType(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewMidacationRepo(db)

	rows := sqlmock.NewRows([]string{"name", "type", "quantity"}).
		AddRow("Med1", "Type1", 10).
		AddRow("Med2", "Type2", 5).
		AddRow("Med3", "Type1", 15)
	mock.ExpectQuery(`SELECT name, type, quantity FROM medications`).
		WillReturnRows(rows)

	expected := &models.MedicarionaGrouped{
		Medications: []models.MedicinesGetAll{
			{Type: "Type1", Count: 2, Medications: []models.MedicationsGet{
				{Name: "Med1", Type: "Type1", Quantity: 10},
				{Name: "Med3", Type: "Type1", Quantity: 15},
			}},
			{Type: "Type2", Count: 1, Medications: []models.MedicationsGet{
				{Name: "Med2", Type: "Type2", Quantity: 5},
			}},
		},
	}

	result, err := repo.GetMedicationsGroupedByType("")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}
