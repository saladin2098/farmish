package managers

import (
	"database/sql"
	"farmish/models"
)

type AnimalRepo struct {
	Conn *sql.DB
}

func NewAnimalRepo(db *sql.DB) *AnimalRepo {
	return &AnimalRepo{Conn: db}
}

func (m *AnimalRepo) GetAllAnimalIds() ([]int, error) {
	query := "SELECT id FROM animals"
	rows, err := m.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	return ids, nil
}

func (m *AnimalRepo) CreateAnimal(animal *models.AnimalCreate, avg_water, avg_consumtion float64, gen_id int) (*models.Animal, error) {
	tx, err := m.Conn.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err

	}
	created_animal := models.Animal{}
	query1 := `
	INSERT INTO animals (id, type, birth, weight, avg_water, avg_consumption)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING *
	`
	err = tx.QueryRow(
		query1, animal.ID, animal.Type, animal.Birth, animal.Weight,
		avg_water, avg_consumtion,
	).Scan(
		&created_animal.ID, &created_animal.Type, &created_animal.Birth,
		&created_animal.Weight, &created_animal.AvgWater, &created_animal.AvgConsumption,
		&created_animal.CreatedAt, &created_animal.UpdatedAt, &created_animal.DeletedAt,
	)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	is_treated := false
	if animal.IsHealthy {
		is_treated = true
		animal.Condition = "Healthy"
		animal.Medication = "None"
	} else {
		is_treated = false
	}
	query2 := `
	INSERT INTO health_conditions (id, animal_id, is_healthy, condition, medication, is_treated)
	VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err = tx.Exec(query2, gen_id, animal.ID, animal.IsHealthy, animal.Condition, animal.Medication, is_treated)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &created_animal, nil
}

func (m *AnimalRepo) UpdateAnimal(animal *models.AnimalUpdate) error {
	tx, err := m.Conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query1 := "UPDATE animals SET weight = $1 WHERE id = $1"
	_, err = tx.Exec(query1, animal.Weight, animal.IsHealthy, animal.Condition, animal.Medication, animal.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	query2 := "UPDATE health_conditions SET is_healthy = $1, condition = $2, medication = $3 WHERE animal_id = $4"
	_, err = tx.Exec(query2, animal.IsHealthy, animal.Condition, animal.Medication, animal.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
