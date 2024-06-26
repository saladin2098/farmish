package managers

import (
	"database/sql"
	"farmish/models"
)

type ScheduleRepo struct {
	DB *sql.DB
}

func NewScheduleRepo(conn *sql.DB) *ScheduleRepo {
	return &ScheduleRepo{DB: conn}
}

func (m *ScheduleRepo) GetAllScheduleIds() ([]int, error) {
	query := "SELECT id FROM schedules"
	rows, err := m.DB.Query(query)
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

func (r *ScheduleRepo) CreateSchedule(s *models.Schedule) (*models.Schedule, error) {
	query := `INSERT INTO schedules (id, time1, time2, time3) VALUES ($1, $2, $3, $4) RETURNING id, time1, time2, time3`
	var createdSchedule models.Schedule
	err := r.DB.QueryRow(query, s.ID, s.Time1, s.Time2, s.Time3).Scan(
		&createdSchedule.ID, &createdSchedule.Time1, &createdSchedule.Time2, &createdSchedule.Time3)
	if err != nil {
		return nil, err
	}
	return &createdSchedule, nil
}
func (r *ScheduleRepo) GetSchedule(id int) (*models.Schedule, error) {
	query := `SELECT id, time1, time2, time3 FROM schedules WHERE id = $1`
	var schedule models.Schedule
	err := r.DB.QueryRow(query, id).Scan(&schedule.ID, &schedule.Time1, &schedule.Time2, &schedule.Time3)
	if err != nil {
		return nil, err
	}
	return &schedule, nil
}

func (r *ScheduleRepo) UpdateSchedule(s *models.Schedule) error {
	query := `UPDATE schedules SET time1 = $1, time2 = $2, time3 = $3 WHERE id = $4`
	_, err := r.DB.Exec(query, s.Time1, s.Time2, s.Time3, s.ID)
	return err
}

func (r *ScheduleRepo) DeleteSchedule(id int) error {
	query := `DELETE FROM schedules WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

type FeedingScheduleRepo struct {
	DB *sql.DB
}

func NewFeedingScheduleRepo(db *sql.DB) *FeedingScheduleRepo {
	return &FeedingScheduleRepo{DB: db}
}

func (r *FeedingScheduleRepo) CreateFeedingSchedule(fs *models.FeedingSchedule) (*models.FeedingSchedule, error) {
	query := `INSERT INTO feeding_schedule (
		id, 
		animal_type, 
		last_fed_index, 
		next_fed_index, 
		schedule_id) VALUES ($1, $2, $3, $4, $5) RETURNING id, animal_type, last_fed_index, next_fed_index, schedule_id`
	var createdFeedingSchedule models.FeedingSchedule
	err := r.DB.QueryRow(query,
		fs.ID,
		fs.AnimalType,
		fs.LastFedIndex,
		fs.NextFedIndex,
		fs.ScheduleID).Scan(
		&createdFeedingSchedule.ID,
		&createdFeedingSchedule.AnimalType,
		&createdFeedingSchedule.LastFedIndex,
		&createdFeedingSchedule.NextFedIndex,
		&createdFeedingSchedule.ScheduleID)
	if err != nil {
		return nil, err
	}
	return &createdFeedingSchedule, nil
}

func (r *FeedingScheduleRepo) GetFeedingSchedule(id int) (*models.FeedingSchedule, error) {
	query := `SELECT id, 
		animal_type, 
		last_fed_index, 
		next_fed_index, 
		schedule_id FROM feeding_schedule WHERE id = $1`
	var feedingSchedule models.FeedingSchedule
	err := r.DB.QueryRow(query, id).Scan(
		&feedingSchedule.ID,
		&feedingSchedule.AnimalType,
		&feedingSchedule.LastFedIndex,
		&feedingSchedule.NextFedIndex,
		&feedingSchedule.ScheduleID)
	if err != nil {
		return nil, err
	}
	return &feedingSchedule, nil
}

func (r *FeedingScheduleRepo) UpdateFeedingSchedule(fs *models.FeedingSchedule) error {
	query := `UPDATE feeding_schedule SET animal_type = $1, last_fed_index = $2, next_fed_index = $3, schedule_id = $4 WHERE id = $5`
	_, err := r.DB.Exec(query, fs.AnimalType, fs.LastFedIndex, fs.NextFedIndex, fs.ScheduleID, fs.ID)
	return err
}

func (r *FeedingScheduleRepo) DeleteFeedingSchedule(id int) error {
	query := `DELETE FROM feeding_schedule WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
