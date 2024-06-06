package postgresql

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
	var existingMed m.Medications
	checkQuery := `SELECT id, name, type, quantity FROM medications WHERE name = $1 AND type = $2`
	err := r.DB.QueryRow(checkQuery, med.Name, med.Type).Scan(
		&existingMed.ID,
		&existingMed.Name,
		&existingMed.Type,
		&existingMed.Quantity)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		insertQuery := `INSERT INTO medications(id,name, type, quantity) VALUES ($1, $2, $3, $4) RETURNING id, name, type, quantity`
		err = r.DB.QueryRow(insertQuery,
			med.ID,
			med.Name,
			med.Type,
			med.Quantity).Scan(
			&existingMed.ID,
			&existingMed.Name,
			&existingMed.Type,
			&existingMed.Quantity)
		if err != nil {
			return nil, err
		}
	} else {
		// Medication exists, update the quantity
		updateQuery := `UPDATE medications SET quantity = quantity + $1 WHERE id = $2 RETURNING id, name, type, quantity`
		err = r.DB.QueryRow(updateQuery,
			med.Quantity,
			existingMed.ID).Scan(
			&existingMed.ID,
			&existingMed.Name,
			&existingMed.Type,
			&existingMed.Quantity)
		if err != nil {
			return nil, err
		}
	}

	return &existingMed, nil
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
	if tur!= "" {
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
			Count: len(medsCopy),
			Medications: medsCopy,
		})
	}
	res := m.MedicarionaGrouped{Medications: result}

	return &res, nil
}
func (r *MedicationRepo) GetAllMedicationIDs() (*[]int,error) {
	var ids []int
    query := `select id from medications`
    rows,err := r.DB.Query(query)
    if err!= nil {
        return nil,err
    }
    defer rows.Close()
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err!= nil {
            return nil,err
        }
        ids = append(ids, id)
    }
    if err := rows.Err(); err!= nil {
        return nil,err
    }
    return &ids,nil
}