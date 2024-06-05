package managers

import (
	"farmish/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestCreateSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewScheduleRepo(db)

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

	repo := NewScheduleRepo(db)

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

	repo := NewScheduleRepo(db)

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

	repo := NewScheduleRepo(db)

	mock.ExpectExec("DELETE FROM schedules WHERE id = \\$1").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteSchedule(1)
	assert.NoError(t, err)
}

func TestCreateFeedingSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewFeedingScheduleRepo(db)

	feedingSchedule := &models.FeedingSchedule{
		ID:           1,
		AnimalType:   "Cow",
		LastFedIndex: 1,
		NextFedIndex: 2,
		ScheduleID:   1,
	}

	mock.ExpectQuery("INSERT INTO feeding_schedules").
		WithArgs(feedingSchedule.ID, feedingSchedule.AnimalType, feedingSchedule.LastFedIndex, feedingSchedule.NextFedIndex, feedingSchedule.ScheduleID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "animal_type", "last_fed_index", "next_fed_index", "schedule_id"}).
			AddRow(feedingSchedule.ID, feedingSchedule.AnimalType, feedingSchedule.LastFedIndex, feedingSchedule.NextFedIndex, feedingSchedule.ScheduleID))

	createdFeedingSchedule, err := repo.CreateFeedingSchedule(feedingSchedule)
	assert.NoError(t, err)
	assert.NotNil(t, createdFeedingSchedule)
	assert.Equal(t, feedingSchedule.ID, createdFeedingSchedule.ID)
}

func TestGetFeedingSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewFeedingScheduleRepo(db)

	expectedFeedingSchedule := &models.FeedingSchedule{
		ID:           1,
		AnimalType:   "Cow",
		LastFedIndex: 1,
		NextFedIndex: 2,
		ScheduleID:   1,
	}

	mock.ExpectQuery("SELECT id, animal_type, last_fed_index, next_fed_index, schedule_id FROM feeding_schedules WHERE id =").
		WithArgs(expectedFeedingSchedule.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "animal_type", "last_fed_index", "next_fed_index", "schedule_id"}).
			AddRow(expectedFeedingSchedule.ID, expectedFeedingSchedule.AnimalType, expectedFeedingSchedule.LastFedIndex, expectedFeedingSchedule.NextFedIndex, expectedFeedingSchedule.ScheduleID))

	feedingSchedule, err := repo.GetFeedingSchedule(expectedFeedingSchedule.ID)
	assert.NoError(t, err)
	assert.NotNil(t, feedingSchedule)
	assert.Equal(t, expectedFeedingSchedule, feedingSchedule)
}

func TestUpdateFeedingSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewFeedingScheduleRepo(db)

	feedingSchedule := &models.FeedingSchedule{
		ID:           1,
		AnimalType:   "Cow",
		LastFedIndex: 1,
		NextFedIndex: 2,
		ScheduleID:   1,
	}

	mock.ExpectExec("UPDATE feeding_schedules SET animal_type = \\$1, last_fed_index = \\$2, next_fed_index = \\$3, schedule_id = \\$4 WHERE id = \\$5").
		WithArgs(feedingSchedule.AnimalType, feedingSchedule.LastFedIndex, feedingSchedule.NextFedIndex, feedingSchedule.ScheduleID, feedingSchedule.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.UpdateFeedingSchedule(feedingSchedule)
	assert.NoError(t, err)
}

func TestDeleteFeedingSchedule(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	repo := NewFeedingScheduleRepo(db)

	mock.ExpectExec("DELETE FROM feeding_schedules WHERE id = \\$1").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.DeleteFeedingSchedule(1)
	assert.NoError(t, err)
}
