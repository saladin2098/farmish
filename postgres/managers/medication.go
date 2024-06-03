package managers

import (
	"database/sql"
	m "farmish/models"
	"fmt"
	"strconv"
	"strings"
)

type MedicationRepo struct {
	DB *sql.DB
}

func NewMidacationRepo(db *sql.DB) *MedicationRepo {
	return &MedicationRepo{DB: db}
}

func (r *MedicationRepo) CreateMedication(med *m.Medications) (*m.Medications, error) {
	query := `insert into medications(id, name, type, quantity) values ($1, $2, $3, $4) returning id, name, type, quantity`

	var createdMed m.Medications
	err := r.DB.QueryRow(query,
		med.ID,
		med.Name,
		med.Type,
		med.Quantity).Scan(
		&createdMed.ID,
		&createdMed.Name,
		&createdMed.Type,
		&createdMed.Quantity)
	if err != nil {
		return nil, err
	}

	return &createdMed, nil
}
func (r *MedicationRepo) GetMedication(id int, name, turi string) (*m.Medications, error) {
	baseQuery := `select id, name, type, quantity from medications where `
	var args []interface{}
	var conditions []string

	if id != 0 {
		conditions = append(conditions, "id = $"+strconv.Itoa(len(args)+1))
		args = append(args, id)
	}
	if name != "" {
		conditions = append(conditions, "name = $"+strconv.Itoa(len(args)+1))
		args = append(args, name)
	}
	if turi != "" {
		conditions = append(conditions, "type = $"+strconv.Itoa(len(args)+1))
		args = append(args, turi)
	}

	var query string
	if len(args) > 1 {
		query = baseQuery + " " + strings.Join(conditions, " and ")
	} else {
		query = baseQuery + " " + conditions[0]
	}

	var med m.Medications
	err := r.DB.QueryRow(query, args...).Scan(&med.ID, &med.Name, &med.Type, &med.Quantity)
	if err != nil {
		return nil, err
	}

	return &med, nil
}
func (r *MedicationRepo) DeleteMedication(id int) error {
	query := `delete from medications where id = $1`
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
func (r *MedicationRepo) UpdateMedication(med *m.Medications) error {
	query := `update medications set name = $1, type = $2, quantity = $3 where id = $4`
	_, err := r.DB.Exec(query, med.Name, med.Type, med.Quantity, med.ID)
	if err != nil {
		return err
	}

	return nil
}
func (r *MedicationRepo) GetMedicationsGroupedByType(tur string) (*m.MedicarionaGrouped, error) {
	query := `
		SELECT name, type, quantity
		FROM medications
	`
	if tur != "" {
		query += fmt.Sprintf(" where type = '%s'", tur)
	}

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	typeMap := make(map[string][]m.MedicationsGet)

	for rows.Next() {
		var med m.MedicationsGet
		if err := rows.Scan(&med.Name, &med.Type, &med.Quantity); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		typeMap[med.Type] = append(typeMap[med.Type], med)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %v", err)
	}

	var result []m.MedicinesGetAll
	for medType, meds := range typeMap {
		medsCopy := meds
		result = append(result, m.MedicinesGetAll{
			Type:        medType,
			Count:       len(medsCopy),
			Medications: medsCopy,
		})
	}
	res := m.MedicarionaGrouped{Medications: result}

	return &res, nil
}
