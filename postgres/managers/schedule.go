package managers

import (
	"database/sql"
	"farmish/models"
)

type ScheduleRepo struct {
	Conn *sql.DB
}

func NewScheduleRepo(conn *sql.DB) *ScheduleRepo {
	return &ScheduleRepo{Conn: conn}
}

func (m *ScheduleRepo) GetAllScheduleIds() ([]int, error) {
	query := "SELECT id FROM schedules"
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

func (m *ScheduleRepo) CreateSchedule(schedule *models.Schedule, feedingSchedule *models.FeedingSchedule) error {
	tx, err := m.Conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	query1 := "INSERT INTO schedules (id, animal_type, time1, time2, time3) VALUES ($1, $2, $3, $4, $5)"
	_, err = tx.Exec(query1, schedule.ID, schedule.AnimalType, schedule.Time1, schedule.Time2, schedule.Time3)
	if err != nil {
		tx.Rollback()
		return err
	}

	query2 := "INSERT INTO feeding_schedules (id, animal_type, last_fed_index) VALUES ($1, $2, $3)"
	_, err = tx.Exec(query2, feedingSchedule.ID, schedule.AnimalType, feedingSchedule.LastFedIndex)
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
