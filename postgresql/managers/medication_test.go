package managers

import (
	"database/sql"
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

	testCases := []struct {
		name         string
		inputMed     *models.Medications
		existingMed  *models.Medications
		expectInsert bool
		expectedMed  *models.Medications
		expectError  bool
	}{
		{
			name: "Create new medication",
			inputMed: &models.Medications{
				ID:       1,
				Name:     "Paracetamol",
				Type:     "Painkiller",
				Quantity: 100,
			},
			existingMed:  nil,
			expectInsert: true,
			expectedMed: &models.Medications{
				ID:       1,
				Name:     "Paracetamol",
				Type:     "Painkiller",
				Quantity: 100,
			},
			expectError: false,
		},
		{
			name: "Update existing medication",
			inputMed: &models.Medications{
				ID:       1,
				Name:     "Paracetamol",
				Type:     "Painkiller",
				Quantity: 50,
			},
			existingMed: &models.Medications{
				ID:       1,
				Name:     "Paracetamol",
				Type:     "Painkiller",
				Quantity: 100,
			},
			expectInsert: false,
			expectedMed: &models.Medications{
				ID:       1,
				Name:     "Paracetamol",
				Type:     "Painkiller",
				Quantity: 150,
			},
			expectError: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.existingMed != nil {
				mock.ExpectQuery(`SELECT id, name, type, quantity FROM medications WHERE name = \$1 AND type = \$2`).
					WithArgs(tc.inputMed.Name, tc.inputMed.Type).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "quantity"}).
						AddRow(tc.existingMed.ID, tc.existingMed.Name, tc.existingMed.Type, tc.existingMed.Quantity))
			} else {
				mock.ExpectQuery(`SELECT id, name, type, quantity FROM medications WHERE name = \$1 AND type = \$2`).
					WithArgs(tc.inputMed.Name, tc.inputMed.Type).
					WillReturnError(sql.ErrNoRows)
			}

			if tc.expectInsert {
				mock.ExpectQuery(`INSERT INTO medications\(id,name, type, quantity\) VALUES \(\$1, \$2, \$3, \$4\) RETURNING id, name, type, quantity`).
					WithArgs(tc.inputMed.ID, tc.inputMed.Name, tc.inputMed.Type, tc.inputMed.Quantity).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "quantity"}).
						AddRow(tc.expectedMed.ID, tc.expectedMed.Name, tc.expectedMed.Type, tc.expectedMed.Quantity))
			} else {
				mock.ExpectQuery(`UPDATE medications SET quantity = quantity \+ \$1 WHERE id = \$2 RETURNING id, name, type, quantity`).
					WithArgs(tc.inputMed.Quantity, tc.existingMed.ID).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "quantity"}).
						AddRow(tc.expectedMed.ID, tc.expectedMed.Name, tc.expectedMed.Type, tc.expectedMed.Quantity))
			}

			actualMed, err := repo.CreateMedication(tc.inputMed)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedMed, actualMed)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
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

	testCases := []struct {
		name        string
		tur         string
		expectedMed *models.MedicarionaGrouped
	}{
		{
			name: "Get all medications",
			tur:  "",
			expectedMed: &models.MedicarionaGrouped{
				Medications: []models.MedicinesGetAll{
					{
						Type:  "Painkiller",
						Count: 2,
						Medications: []models.MedicationsGet{
							{Name: "Paracetamol", Type: "Painkiller", Quantity: 100},
							{Name: "Ibuprofen", Type: "Painkiller", Quantity: 200},
						},
					},
					{
						Type:  "Antibiotic",
						Count: 1,
						Medications: []models.MedicationsGet{
							{Name: "Amoxicillin", Type: "Antibiotic", Quantity: 50},
						},
					},
				},
			},
		},
		{
			name: "Get medications by type",
			tur:  "Painkiller",
			expectedMed: &models.MedicarionaGrouped{
				Medications: []models.MedicinesGetAll{
					{
						Type:  "Painkiller",
						Count: 2,
						Medications: []models.MedicationsGet{
							{Name: "Paracetamol", Type: "Painkiller", Quantity: 100},
							{Name: "Ibuprofen", Type: "Painkiller", Quantity: 200},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			query := `SELECT name, type, quantity FROM medications`
			if tc.tur != "" {
				query += ` where type = '` + tc.tur + `'`
			}

			rows := sqlmock.NewRows([]string{"name", "type", "quantity"})
			if tc.tur == "" || tc.tur == "Painkiller" {
				rows.AddRow("Paracetamol", "Painkiller", 100)
				rows.AddRow("Ibuprofen", "Painkiller", 200)
			}
			if tc.tur == "" || tc.tur == "Antibiotic" {
				rows.AddRow("Amoxicillin", "Antibiotic", 50)
			}

			mock.ExpectQuery(query).WillReturnRows(rows)

			actualMed, err := repo.GetMedicationsGroupedByType(tc.tur)

			assert.NoError(t, err)
			assert.Equal(t, tc.expectedMed, actualMed)
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
