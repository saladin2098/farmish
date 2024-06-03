package managers

import (
	"farmish/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateAnimal(t *testing.T) {
	mock, repo := testSetupAnimal(t)

	animal := &models.AnimalCreate{
		ID:         1,
		Type:       "Cow",
		Birth:      "12-12-2012",
		Weight:     200,
		IsHealthy:  true,
		Condition:  "Healthy",
		Medication: "x",
	}

	rows := sqlmock.NewRows([]string{"id", "type", "birth", "weight", "is_healthy", "condition", "medication"}).
		AddRow(animal.ID, animal.Type, animal.Birth, animal.Weight, animal.IsHealthy, animal.Condition, animal.Medication)
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO animals`).
		WithArgs(animal.ID, animal.Type, animal.Birth, animal.Weight, animal.IsHealthy, animal.Condition, animal.Medication).
		WillReturnRows(rows)
	mock.ExpectExec(`INSERT INTO health_conditions`).
		WithArgs(sqlmock.AnyArg(), animal.ID, animal.IsHealthy, animal.Condition, animal.Medication, animal.IsHealthy, false).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	createdAnimal, err := repo.CreateAnimal(animal, 0, 0, 1)
	assert.NoError(t, err)
	assert.Equal(t, animal.ID, createdAnimal.ID)
}

func TestUpdateAnimal(t *testing.T) {
	mock, repo := testSetupAnimal(t)

	animal := &models.AnimalUpdate{
		ID:         1,
		Weight:     200,
		IsHealthy:  true,
		Condition:  "Healthy",
		Medication: "x",
	}

	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE animals`).
		WithArgs(animal.Weight, animal.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(`UPDATE health_conditions`).
		WithArgs(animal.IsHealthy, animal.Condition, animal.Medication, animal.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.UpdateAnimal(animal)
	assert.NoError(t, err)
}

func testSetupAnimal(t *testing.T) (sqlmock.Sqlmock, *AnimalRepo) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := NewAnimalRepo(db)

	return mock, repo
}
