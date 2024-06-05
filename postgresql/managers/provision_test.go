package managers

import (
	"farmish/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAnimalRepo_GetAnimalByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAnimalRepo(db)

	rows := sqlmock.NewRows([]string{"id", "type", "birth", "weight", "avg_consumption", "avg_water", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Cow", "2020-01-01", 500, 30.0, 40.0, time.Now(), time.Now(), 0)

	mock.ExpectQuery("SELECT id, type, birth, weight, avg_consumption, avg_water, created_at, updated_at, deleted_at FROM animals WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	animal, err := repo.GetAnimalByID(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedAnimal := &models.Animal{
		ID:             1,
		Type:           "Cow",
		Birth:          "2020-01-01",
		Weight:         500,
		AvgConsumption: 30.0,
		AvgWater:       40.0,
	}

	if animal.ID != expectedAnimal.ID {
		t.Errorf("expected ID %v, got %v", expectedAnimal.ID, animal.ID)
	}
}

func TestAnimalRepo_GetAnimalAgeInDays(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAnimalRepo(db)

	rows := sqlmock.NewRows([]string{"age_in_days"}).AddRow(365)

	mock.ExpectQuery("SELECT now\\(\\) - birth AS age_in_days FROM animals WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	ageInDays, err := repo.GetAnimalAgeInDays(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if ageInDays != 365 {
		t.Errorf("expected age in days %v, got %v", 365, ageInDays)
	}
}

func TestAnimalRepo_GetAllAnimalIds(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAnimalRepo(db)

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2)

	mock.ExpectQuery("SELECT id FROM animals").WillReturnRows(rows)

	ids, err := repo.GetAllAnimalIds()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(ids) != 2 {
		t.Errorf("expected 2 animal IDs, got %v", len(ids))
	}
}

func TestAnimalRepo_CreateAnimal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAnimalRepo(db)

	newAnimal := &models.AnimalCreate{
		ID:         2,
		Type:       "Chicken",
		Birth:      "2023-01-01",
		Weight:     3,
		AnimalType: "Poultry",
		IsHealthy:  true,
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO animals").WithArgs(newAnimal.ID, newAnimal.Type, newAnimal.AnimalType, newAnimal.Birth, newAnimal.Weight, 10.0, 5.0).
		WillReturnRows(sqlmock.NewRows([]string{"id", "type", "animal_type", "birth", "weight", "avg_water", "avg_consumption", "created_at", "updated_at", "deleted_at"}).
			AddRow(newAnimal.ID, newAnimal.Type, newAnimal.AnimalType, newAnimal.Birth, newAnimal.Weight, 10.0, 5.0, time.Now(), time.Now(), 0))
	mock.ExpectExec("INSERT INTO health_conditions").WithArgs(1, newAnimal.ID, newAnimal.IsHealthy, "Healthy", "None").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	createdAnimal, err := repo.CreateAnimal(newAnimal, 10.0, 5.0, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if createdAnimal.ID != newAnimal.ID {
		t.Errorf("expected ID %v, got %v", newAnimal.ID, createdAnimal.ID)
	}
}

func TestAnimalRepo_UpdateAnimal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAnimalRepo(db)

	updateAnimal := &models.AnimalUpdate{
		ID:         1,
		Weight:     550,
		IsHealthy:  true,
		Condition:  "Healthy",
		Medication: "None",
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE animals SET weight = \\$1 WHERE id = \\$2").WithArgs(updateAnimal.Weight, updateAnimal.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("UPDATE health_conditions SET is_healthy = \\$1, condition = \\$2, medication = \\$3 WHERE animal_id = \\$4").
		WithArgs(updateAnimal.IsHealthy, updateAnimal.Condition, updateAnimal.Medication, updateAnimal.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.UpdateAnimal(updateAnimal)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestAnimalRepo_DeleteAnimal(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewAnimalRepo(db)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE animals SET deleted_at = EXTRACT\\(EPOCH FROM NOW\\(\\)\\) WHERE id = \\$1").
		WithArgs(2).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("DELETE FROM health_conditions WHERE animal_id = \\$1").
		WithArgs(2).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err = repo.DeleteAnimal(2)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
