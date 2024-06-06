package managers

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestFeedAnimals(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock DB: %s", err)
	}
	defer db.Close()

	repo := NewFeedingRepo(db)

	TimeNow = func() time.Time {
		return time.Date(2023, 10, 23, 10, 15, 0, 0, time.UTC)
	}

	mock.ExpectBegin()

	mock.ExpectQuery(`SELECT schedule_id, last_fed_index, next_fed_index FROM feeding_schedules WHERE animal_type = \$1`).
		WithArgs("ot").
		WillReturnRows(sqlmock.NewRows([]string{"schedule_id", "last_fed_index", "next_fed_index"}).AddRow(20878, 3, 1))

	mock.ExpectQuery(`SELECT EXTRACT\(HOUR FROM time1\) AS hour1, EXTRACT\(HOUR FROM time2\) AS hour2, EXTRACT\(HOUR FROM time3\) AS hour3 FROM schedules WHERE id = \$1`).
		WithArgs(20878).
		WillReturnRows(sqlmock.NewRows([]string{"hour1", "hour2", "hour3"}).AddRow(10, 13, 19))

	mock.ExpectQuery(`SELECT animal_type, quantity FROM provision WHERE type = \$1`).
		WithArgs("provisionType").
		WillReturnRows(sqlmock.NewRows([]string{"animal_type", "quantity"}).AddRow("ot", 50.0))

	mock.ExpectQuery(`SELECT total FROM water_consumption`).
		WillReturnRows(sqlmock.NewRows([]string{"total"}).AddRow(100.0))

	mock.ExpectQuery(`SELECT avg_consumption, avg_water FROM animals WHERE type = \$1`).
		WithArgs("ot").
		WillReturnRows(sqlmock.NewRows([]string{"avg_consumption", "avg_water"}).AddRow(10.0, 5.0))

	mock.ExpectExec(`UPDATE provision SET quantity = \$1 WHERE type = \$2`).
		WithArgs(40.0, "provisionType").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`UPDATE water_consumption SET total = \$1`).
		WithArgs(90.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(`UPDATE feeding_schedules SET last_fed_index = \$1, next_fed_index = \$2 WHERE animal_type = \$3`).
		WithArgs(1, 2, "ot").
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err = repo.FeedAnimals("ot", "provisionType")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
