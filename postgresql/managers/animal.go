package managers

import (
	"database/sql"
	"farmish/models"
	"fmt"
)

type AnimalRepo struct {
	Conn *sql.DB
}

func NewAnimalRepo(db *sql.DB) *AnimalRepo {
	return &AnimalRepo{Conn: db}
}

func (m *AnimalRepo) GetAnimalByID(id int) (*models.Animal, error) {
	query := "SELECT id, type, birth, weight, avg_consumption, avg_water, created_at, updated_at, deleted_at FROM animals WHERE id = $1"
	animal := models.Animal{}
	row := m.Conn.QueryRow(query, id)
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

func (m *AnimalRepo) GetAllAnimals(animal_type string, is_healthy, is_hungry string) (*models.AnimalsGetAll, error) {
	all_animals := models.AnimalsGetAll{}
	query := `
	SELECT a.id, a.type, a.animal_type, a.birth, a.weight, 
	h.is_healthy, h.condition, h.medication, 
	a.avg_consumption, a.avg_water
	FROM animals a
	LEFT JOIN health_conditions h ON a.id = h.animal_id
	LEFT JOIN feeding_schedules fs ON a.animal_type = fs.animal_type
	LEFT JOIN schedules s ON fs.schedule_id = s.id
	WHERE a.deleted_at = 0
	`
	var agrs []interface{}
	paramIndex := 1
	if animal_type != "" {
		query += fmt.Sprintf(" AND a.animal_type = $%d", paramIndex)
		agrs = append(agrs, animal_type)
		paramIndex++
	}
	if is_healthy == "true" {
		query += fmt.Sprintf(" AND h.is_healthy = $%d", paramIndex)
		agrs = append(agrs, true)
		paramIndex++
	} else if is_healthy == "false" {
		query += fmt.Sprintf(" AND h.is_healthy = $%d", paramIndex)
		agrs = append(agrs, false)
		paramIndex++
	}
	if is_hungry == "true" {
		query += `
		AND (
			(fs.next_fed_index = '1' AND CAST(NOW() AS time) > s.time1) OR
			(fs.next_fed_index = '2' AND CAST(NOW() AS time) > s.time2) OR
			(fs.next_fed_index = '3' AND CAST(NOW() AS time) > s.time3)
		);
		`
	}
	rows, err := m.Conn.Query(query, agrs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		animal := models.AnimalGet{}
		err := rows.Scan(&animal.ID, &animal.Type, &animal.AnimalType, &animal.Birth, &animal.Weight,
			&animal.HealthCondition.IsHealthy, &animal.HealthCondition.Condition, &animal.HealthCondition.Medication,
			&animal.AvgConsumption, &animal.AvgWater,
		)
		if err != nil {
			return nil, err
		}
		all_animals.Animals = append(all_animals.Animals, animal)
		all_animals.Count++
	}
	// fmt.Println(all_animals.Count, "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	return &all_animals, nil
}

func (m *AnimalRepo) CreateAnimal(animal *models.AnimalCreate, avg_water, avg_consumtion float64, gen_id int) (*models.Animal, error) {
	tx, err := m.Conn.Begin()
	if err != nil {
		tx.Rollback()
		return nil, err

	}
	created_animal := models.Animal{}
	query1 := `
	INSERT INTO animals (id, type, animal_type, birth, weight, avg_water, avg_consumption)
	VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *
	`
	err = tx.QueryRow(
		query1, animal.ID, animal.Type, animal.AnimalType, animal.Birth, animal.Weight,
		avg_water, avg_consumtion,
	).Scan(
		&created_animal.ID, &created_animal.Type, &created_animal.AnimalType,
		&created_animal.Birth, &created_animal.Weight, &created_animal.AvgWater, &created_animal.AvgConsumption,
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
		return err
	}
	query1 := "UPDATE animals SET weight = $1 WHERE id = $2"
	if _, err := tx.Exec(query1, animal.Weight, animal.ID); err != nil {
		tx.Rollback()
		return err
	}
	query2 := "UPDATE health_conditions SET is_healthy = $1, condition = $2, medication = $3 WHERE animal_id = $4"
	if _, err := tx.Exec(query2, animal.IsHealthy, animal.Condition, animal.Medication, animal.ID); err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
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
		return err
	}
	query1 := "UPDATE animals SET deleted_at = EXTRACT(EPOCH FROM NOW()) WHERE id = $1"
	if _, err := m.Conn.Exec(query1, id); err != nil {
		return err
	}
	query2 := "DELETE FROM health_conditions WHERE animal_id = $1"
	if _, err = m.Conn.Exec(query2, id); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (m *AnimalRepo) GetNextFedIndex(animal_type string) (int, error) {
	query := "SELECT next_fed_index FROM feeding_schedules WHERE animal_type = $1"
	row := m.Conn.QueryRow(query, animal_type)
	var nextFedIndex int
	err := row.Scan(&nextFedIndex)
	if err != nil {
		return 0, err
	}
	return nextFedIndex, nil
}

func (m *AnimalRepo) GetLastFedTime(index int, schedule_id int) (string, error) {
	query := ""
	if index == 1 {
		query = "SELECT time1 FROM schedules WHERE id = $1"

	} else if index == 2 {
		query = "SELECT time2 FROM schedules WHERE id = $1"
	} else if index == 3 {
		query = "SELECT time3 FROM schedules WHERE id = $1"
	}
	row := m.Conn.QueryRow(query, schedule_id)
	var lastFedTime string
	err := row.Scan(&lastFedTime)
	if err != nil {
		return "", err
	}
	return lastFedTime, nil
}

func (m *AnimalRepo) GetScheduleID(animal_type string) (int, error) {
	query := "SELECT schedule_id FROM feeding_schedules WHERE animal_type = $1"
	row := m.Conn.QueryRow(query, animal_type)
	var scheduleID int
	err := row.Scan(&scheduleID)
	if err != nil {
		return 0, err
	}
	return scheduleID, nil
}
