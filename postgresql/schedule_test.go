package postgresql_test

import (
	"farmish/models"
	"farmish/postgresql"
	"strconv"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgresql.NewScheduleRepo(db)

	schedule := &models.Schedule{
		ID:    1,
		Time1: time.Now(),
		Time2: time.Now().Add(2 * time.Hour),
		Time3: time.Now().Add(4 * time.Hour),
	}

	mock.ExpectQuery("INSERT INTO schedules").
		WithArgs(schedule.ID, schedule.Time1, schedule.Time2, schedule.Time3).
		WillReturnRows(sqlmock.NewRows([]string{"id", "time1", "time2", "time3"}).AddRow(schedule.ID, schedule.Time1, schedule.Time2, schedule.Time3))

	createdSchedule, err := repo.CreateSchedule(schedule)
	assert.NoError(t, err)
	assert.NotNil(t, createdSchedule)
	assert.Equal(t, schedule.ID, createdSchedule.ID)
}

func TestGetSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgresql.NewScheduleRepo(db)

	expectedSchedule := &models.Schedule{
		ID:    1,
		Time1: time.Now(),
		Time2: time.Now().Add(2 * time.Hour),
		Time3: time.Now().Add(4 * time.Hour),
	}

	mock.ExpectQuery("SELECT id, time1, time2, time3 FROM schedules WHERE id =").
		WithArgs(expectedSchedule.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "time1", "time2", "time3"}).AddRow(expectedSchedule.ID, expectedSchedule.Time1, expectedSchedule.Time2, expectedSchedule.Time3))

	schedule, err := repo.GetSchedule(expectedSchedule.ID)
	assert.NoError(t, err)
	assert.NotNil(t, schedule)
	assert.Equal(t, expectedSchedule, schedule)
}

func TestUpdateSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgresql.NewScheduleRepo(db)

	schedule := &models.Schedule{
		ID:    1,
		Time1: time.Now(),
		Time2: time.Now().Add(2 * time.Hour),
		Time3: time.Now().Add(4 * time.Hour),
	}

	mock.ExpectExec("UPDATE schedules SET time1 = \\$1, time2 = \\$2, time3 = \\$3 WHERE id = \\$4").
		WithArgs(schedule.Time1, schedule.Time2, schedule.Time3, schedule.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateSchedule(schedule)
	assert.NoError(t, err)
}

func TestDeleteSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgresql.NewScheduleRepo(db)

	mock.ExpectExec("DELETE FROM schedules WHERE id = \\$1").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteSchedule(1)
	assert.NoError(t, err)
}

func TestCreateFeedingSchedule(t *testing.T) {
	// Create mock DB and FeedingScheduleRepo
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgresql.NewFeedingScheduleRepo(db)

	// Define the input and expected output
	inputFS := &models.FeedingSchedule{
		ID:           1,
		AnimalType:   "ot",
		LastFedIndex: 3,
		NextFedIndex: 1,
		ScheduleID:   20878,
	}

	expectedFS := &models.FeedingSchedule{
		ID:           1,
		AnimalType:   "ot",
		LastFedIndex: 3,
		NextFedIndex: 1,
		ScheduleID:   20878,
	}

	// Mock the expected query and result
	mock.ExpectQuery(`INSERT INTO feeding_schedule \(
		id, 
		animal_type, 
		last_fed_index, 
		next_fed_index, 
		schedule_id\) VALUES \(\$1, \$2, \$3, \$4, \$5\) RETURNING id, animal_type, last_fed_index, next_fed_index, schedule_id`).
		WithArgs(inputFS.ID, inputFS.AnimalType, strconv.Itoa(inputFS.LastFedIndex), strconv.Itoa(inputFS.NextFedIndex), inputFS.ScheduleID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "animal_type", "last_fed_index", "next_fed_index", "schedule_id"}).
			AddRow(expectedFS.ID, expectedFS.AnimalType, expectedFS.LastFedIndex, expectedFS.NextFedIndex, expectedFS.ScheduleID))

	// Call the function
	actualFS, err := repo.CreateFeedingSchedule(inputFS)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedFS, actualFS)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetFeedingSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgresql.NewFeedingScheduleRepo(db)

	inputID := 1
	expectedFS := &models.FeedingSchedule{
		ID:           1,
		AnimalType:   "ot",
		LastFedIndex: 3,
		NextFedIndex: 1,
		ScheduleID:   20878,
	}

	mock.ExpectQuery(`SELECT id, animal_type, last_fed_index, next_fed_index, schedule_id FROM feeding_schedule WHERE id = \$1`).
		WithArgs(inputID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "animal_type", "last_fed_index", "next_fed_index", "schedule_id"}).
			AddRow(expectedFS.ID, expectedFS.AnimalType, expectedFS.LastFedIndex, expectedFS.NextFedIndex, expectedFS.ScheduleID))

	actualFS, err := repo.GetFeedingSchedule(inputID)

	assert.NoError(t, err)
	assert.Equal(t, expectedFS, actualFS)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateFeedingSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgresql.NewFeedingScheduleRepo(db)

	inputFS := &models.FeedingSchedule{
		ID:           1,
		AnimalType:   "ot",
		LastFedIndex: 3,
		NextFedIndex: 1,
		ScheduleID:   20878,
	}

	mock.ExpectExec(`UPDATE feeding_schedule SET animal_type = \$1, last_fed_index = \$2, next_fed_index = \$3, schedule_id = \$4 WHERE id = \$5`).
		WithArgs(inputFS.AnimalType, strconv.Itoa(inputFS.LastFedIndex), strconv.Itoa(inputFS.NextFedIndex), inputFS.ScheduleID, inputFS.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateFeedingSchedule(inputFS)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}


func TestDeleteFeedingSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := postgresql.NewFeedingScheduleRepo(db)

	inputID := 1

	mock.ExpectExec(`DELETE FROM feeding_schedule WHERE id = \$1`).
		WithArgs(inputID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteFeedingSchedule(inputID)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
