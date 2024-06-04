package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
	// "time"
)

type FeedingRepo struct {
	DB *sql.DB
}

func NewFeedingRepo(db *sql.DB) *FeedingRepo {
	return &FeedingRepo{DB: db}
}

func (r *FeedingRepo) FeedAnimals(animal string) error {
	currentTime := time.Now()

	// Query the feeding schedule for the given animal
	var scheduleID, lastFedIndex, nextFedIndex int
	err := r.DB.QueryRow(`SELECT 
		schedule_id, 
		last_fed_index, 
		next_fed_index 
		FROM feeding_schedule 
		WHERE animal_type = $1`, animal).Scan(&scheduleID, &lastFedIndex, &nextFedIndex)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("feeding schedule not found for the given animal")
		}
		return err
	}

	// Query the schedule times
	var time1, time2, time3 int
	err = r.DB.QueryRow(`SELECT   
		EXTRACT(HOUR FROM time1) AS hour1, 
		EXTRACT(HOUR FROM time2) AS hour2, 
		EXTRACT(HOUR FROM time3) AS hour3 
		FROM schedules 
		WHERE id = $1`, scheduleID).Scan(&time1, &time2, &time3)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("schedule not found")
		}
		return err
	}

	// Define the feeding times in hours
	feedingTimes := []int{time1, time2, time3}

	// Check if the current time falls within the allowed feeding window
	currentHour := currentTime.Hour()
	var nextFeedingTime int
	allowed := false
	for i, feedingHour := range feedingTimes {
		feedingWindowStart := (feedingHour - 1 + 24) % 24
		if currentHour >= feedingWindowStart && currentHour < feedingHour {
			if lastFedIndex == i+1 {
				return errors.New("animals are already fed")
			}
			// Update the last and next feed indices
			lastFedIndex = i + 1
			nextFedIndex = (i+1)%len(feedingTimes) + 1
			allowed = true
			break
		}
		if feedingHour > currentHour {
			nextFeedingTime = feedingHour
			break
		}
	}

	if !allowed {
		nextFeedingWindowStart := (nextFeedingTime - 1 + 24) % 24
		return fmt.Errorf("next schedule is at %02d:00. You can start feeding animals from %02d:00",
			nextFeedingTime, nextFeedingWindowStart)
	}

	// Update the feeding schedule in the database
	_, err = r.DB.Exec(`UPDATE feeding_schedule SET last_fed_index = $1, next_fed_index = $2 WHERE animal_type = $3`, lastFedIndex, nextFedIndex, animal)
	if err != nil {
		return err
	}

	return nil
}