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

func (m *AnimalRepo) GetAnimalByID(id int) (*models.Animal, error) {
	query := "SELECT id, type, birth, weight, avg_consumption, avg_water, created_at, updated_at, deleted_at FROM animals WHERE id = $1"
	row := m.Conn.QueryRow(query, id)
	animal := models.Animal{}
	err := row.Scan(
		&animal.ID, &animal.Type, &animal.Birth,
		&animal.Weight, &animal.AvgConsumption, &animal.AvgWater,
		&animal.CreatedAt, &animal.UpdatedAt, &animal.DeletedAt,
	)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}

func (m *AnimalRepo) GetAnimalAgeInDays(id int) (int, error) {
	query := "SELECT now() - birth AS age_in_days FROM animals WHERE id = $1"
	row := m.Conn.QueryRow(query, id)
	var ageInDays int
	err := row.Scan(&ageInDays)
	if err != nil {
		return 0, err
	}
	return ageInDays, nil
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
	if animal.IsHealthy {
		animal.Condition = "Healthy"
		animal.Medication = "None"
	}
	query2 := `
	INSERT INTO health_conditions (id, animal_id, is_healthy, condition, medication)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err = tx.Exec(query2, gen_id, animal.ID, animal.IsHealthy, animal.Condition, animal.Medication)
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
	query1 := "UPDATE animals SET weight = $1 WHERE id = $2"
	_, err = tx.Exec(query1, animal.Weight, animal.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if animal.IsHealthy {
		animal.Condition = "Healthy"
		animal.Medication = "None"
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

func (m *AnimalRepo) UpdateAvgConsumption(id int, water, meal float64) error {
	query := "UPDATE animals SET avg_consumption = $1, avg_water = $2 WHERE id = $3"
	_, err := m.Conn.Exec(query, meal, water, id)
	if err != nil {
		return err
	}
	return nil
}

func (m *AnimalRepo) DeleteAnimal(id int) error {
	tx, err := m.Conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query1 := "UPDATE animals SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1"
	_, err = m.Conn.Exec(query1, id)
	if err != nil {
		return err
	}
	query2 := "DELETE FROM health_conditions WHERE animal_id = $1"
	_, err = m.Conn.Exec(query2, id)
	if err != nil {
		return err
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
