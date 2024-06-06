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
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := &FeedingRepo{DB: db}

	// Current time mock
	now := time.Date(2024, time.June, 4, 7, 0, 0, 0, time.UTC)
	timeNow = func() time.Time { return now }

	// Feeding schedule mock
	mock.ExpectBegin()
	mock.ExpectQuery("SELECT schedule_id, last_fed_index, next_fed_index FROM feeding_schedule WHERE animal_type = \\$1").
		WithArgs("cow").
		WillReturnRows(sqlmock.NewRows([]string{"schedule_id", "last_fed_index", "next_fed_index"}).
			AddRow(1, 0, 1)) // Adjusted last_fed_index to not simulate "already fed"

	// Schedules mock
	mock.ExpectQuery("SELECT EXTRACT\\(HOUR FROM time1\\) AS hour1, EXTRACT\\(HOUR FROM time2\\) AS hour2, EXTRACT\\(HOUR FROM time3\\) AS hour3 FROM schedules WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"hour1", "hour2", "hour3"}).
			AddRow(6, 12, 18))

	// Provision mock
	mock.ExpectQuery("SELECT animal_type, quantity FROM provision WHERE type = \\$1").
		WithArgs("hay").
		WillReturnRows(sqlmock.NewRows([]string{"animal_type", "quantity"}).
			AddRow("cow", 100))

	// Animals mock
	mock.ExpectQuery("SELECT avg_consumption, avg_water FROM animals WHERE type = \\$1").
		WithArgs("cow").
		WillReturnRows(sqlmock.NewRows([]string{"avg_consumption", "avg_water"}).
			AddRow(50, 10).
			AddRow(50, 10))

	// Update provision mock
	mock.ExpectExec("UPDATE provision SET quantity = \\$1 WHERE type = \\$2").
		WithArgs(0.0, "hay").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Update water consumption mock
	mock.ExpectExec("UPDATE water_consumption SET total = total + \\$1").
		WithArgs(20.0).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Update feeding schedule mock
	mock.ExpectExec("UPDATE feeding_schedule SET last_fed_index = \\$1, next_fed_index = \\$2 WHERE animal_type = \\$3").
		WithArgs(1, 2, "cow").
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Run the function
	err = repo.FeedAnimals("cow", "hay")
	assert.NoError(t, err)

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
