package managers

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

type FeedingRepo struct {
	DB *sql.DB
}

var timeNow = time.Now

func NewFeedingRepo(db *sql.DB) *FeedingRepo {
	return &FeedingRepo{DB: db}
}
func (r *FeedingRepo) GetAllFeedingSheduleIDs() (*[]int, error) {
	var ids []int
	query := `SELECT id FROM feeding_schedule`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return &ids, nil
}
func (r *FeedingRepo) FeedAnimals(animal string, provision string) error {
	currentTime := timeNow()
	tr, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tr.Rollback()
			panic(p) // re-throw panic after Rollback
		} else if err != nil {
			tr.Rollback()
		} else {
			err = tr.Commit()
		}
	}()

	var scheduleID, lastFedIndex, nextFedIndex int
	err = tr.QueryRow(`SELECT schedule_id, last_fed_index, next_fed_index FROM feeding_schedule WHERE animal_type = $1`, animal).Scan(&scheduleID, &lastFedIndex, &nextFedIndex)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("feeding schedule not found for the given animal")
		}
		return err
	}

	var time1, time2, time3 int
	err = tr.QueryRow(`SELECT EXTRACT(HOUR FROM time1) AS hour1, EXTRACT(HOUR FROM time2) AS hour2, EXTRACT(HOUR FROM time3) AS hour3 FROM schedules WHERE id = $1`, scheduleID).Scan(&time1, &time2, &time3)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("schedule not found")
		}
		return err
	}

	feedingTimes := []int{time1, time2, time3}
	currentHour := currentTime.Hour()
	allowed := false
	nextFeedingTime := -1

	for i, feedingHour := range feedingTimes {
		feedingWindowStart := (feedingHour - 1 + 24) % 24
		feedingWindowEnd := (feedingHour + 2) % 24

		if (feedingWindowStart <= currentHour && currentHour <= feedingWindowEnd) ||
			(feedingWindowStart > feedingWindowEnd && (currentHour >= feedingWindowStart || currentHour <= feedingWindowEnd)) {
			if lastFedIndex == i+1 {
				return errors.New("animals are already fed")
			}
			lastFedIndex = i + 1
			nextFedIndex = (i+1)%len(feedingTimes) + 1
			allowed = true
			break
		}

		if nextFeedingTime == -1 || (feedingHour > currentHour && feedingHour < nextFeedingTime) || (nextFeedingTime < currentHour && feedingHour < nextFeedingTime) {
			nextFeedingTime = feedingHour
		}
	}

	if !allowed {
		nextFeedingWindowStart := (nextFeedingTime - 1 + 24) % 24
		return fmt.Errorf("next schedule is at %02d:00. You can start feeding animals from %02d:00",
			nextFeedingTime, nextFeedingWindowStart)
	}

	var provisionAnType string
	var provisionQuantity float64
	err = tr.QueryRow(`SELECT animal_type, quantity FROM provision WHERE type = $1`, provision).Scan(&provisionAnType, &provisionQuantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("animal types in provision not found")
		}
		return err
	}

	provisionAnimalTypes := extractWords(provisionAnType)
	suitable := false
	for _, animalType := range provisionAnimalTypes {
		if animalType == animal {
			suitable = true
			break
		}
	}
	if !suitable {
		return errors.New("you cannot feed this provision to that animal")
	}
	var existingWater float64
	err = tr.QueryRow(`SELECT total FROM water_consumption`).Scan(&existingWater)
	if err != nil {
		return err
	}

	var totalFoodConsumption, totalWaterConsumption float64
	rows, err := tr.Query(`SELECT avg_consumption, avg_water FROM animals WHERE type = $1`, animal)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var avgConsumption, avgWater float64
		if err := rows.Scan(&avgConsumption, &avgWater); err != nil {
			return err
		}
		totalFoodConsumption += avgConsumption
		totalWaterConsumption += avgWater
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if provisionQuantity < totalFoodConsumption {
		return errors.New("not enough provision quantity")
	}
	if existingWater < totalWaterConsumption {
		return errors.New("not enough water")
	}

	newProvisionQuantity := provisionQuantity - totalFoodConsumption
	_, err = tr.Exec(`UPDATE provision SET quantity = $1 WHERE type = $2`, newProvisionQuantity, provision)
	if err != nil {
		return err
	}
	newWaterLevel := existingWater - totalFoodConsumption
	_, err = tr.Exec(`UPDATE water_consumption SET total = $1`, newWaterLevel)
	if err != nil {
		return err
	}

	_, err = tr.Exec(`UPDATE feeding_schedule SET last_fed_index = $1, next_fed_index = $2 WHERE animal_type = $3`, lastFedIndex, nextFedIndex, animal)
	if err != nil {
		return err
	}

	return nil
}
func extractWords(str string) []string {
	return strings.FieldsFunc(str, func(r rune) bool {
		return r == ',' || r == ' '
	})
}
