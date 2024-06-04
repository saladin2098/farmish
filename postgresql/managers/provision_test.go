package managers

import (
	"farmish/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateProvision(t *testing.T) {
	mock, repo, prs := testSetup(t)

	rows := sqlmock.NewRows([]string{"id", "type", "animal_type", "quantity"}).
		AddRow(prs.ID, prs.Type, prs.AnimalType, prs.Quantity)
	mock.ExpectQuery(`insert into provision`).
		WithArgs(prs.ID, prs.Type, prs.AnimalType, prs.Quantity).
		WillReturnRows(rows)

	createdProvision, err := repo.CreateProvision(*prs)
	assert.NoError(t, err)
	assert.Equal(t, prs, createdProvision)
}

func TestGetProvision(t *testing.T) {
	mock, repo, prs := testSetup(t)

	rows := sqlmock.NewRows([]string{"id", "type", "animal_type", "quantity"}).
		AddRow(prs.ID, prs.Type, prs.AnimalType, prs.Quantity)
	mock.ExpectQuery(`select id, type, animal_type, quantity from provision where id = \$1`).
		WithArgs(prs.ID).
		WillReturnRows(rows)

	retrievedProvision, err := repo.GetProvision(prs.ID, "", "", 0)
	assert.NoError(t, err)
	assert.Equal(t, prs, retrievedProvision)

	mock.ExpectQuery(`select id, type, animal_type, quantity from provision where type = \$1`).
		WithArgs(prs.ID).
		WillReturnRows(rows)

	retrievedProvision, err = repo.GetProvision(0, prs.Type, "", 0)
	assert.NoError(t, err)
	assert.Equal(t, prs, retrievedProvision)

	mock.ExpectQuery(`select id, type, animal_type, quantity from provision where type = \$1, quantity <= \$2`).
		WithArgs(prs.ID).
		WillReturnRows(rows)

	retrievedProvision, err = repo.GetProvision(0, "", prs.AnimalType, prs.Quantity)
	assert.NoError(t, err)
	assert.Equal(t, prs, retrievedProvision)
}

func TestDeleteProvision(t *testing.T) {
	mock, repo, _ := testSetup(t)

	now := time.Now().Unix()

	mock.ExpectExec(`update provision set deleted_at=\$1 where id = \$2`).
		WithArgs(now, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.DeleteProvision(1)
	assert.NoError(t, err)
}

func TestUpdateProvision(t *testing.T) {
	mock, repo, _ := testSetup(t)

	prs := models.UpdateProvision{ID: 1, Type: "Qoy", Quantity: 10}

	mock.ExpectExec(`update provision set type = \$1, quantity = \$2 where id = \$3`).
		WithArgs(prs.ID, prs.Type, prs.Quantity).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.UpdateProvision(&prs)
	assert.NoError(t, err)
}

func testSetup(t *testing.T) (sqlmock.Sqlmock, *ProvisionRepo, *models.CreateProvision) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	repo := NewProvisionRepo(db)

	prs := &models.CreateProvision{
		ID:         1,
		Type:       "Cow",
		AnimalType: "Mammals",
		Quantity:   10,
	}

	return mock, repo, prs
}
