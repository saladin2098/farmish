package postgresql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	// "time"
)

type FeedingRepo struct {
	DB *sql.DB
}

func NewFeedingRepo(db *sql.DB) *FeedingRepo {
	return &FeedingRepo{DB: db}
}

func (r *FeedingRepo) FeedAnimals(animal string, provision string) error {
    currentTime := time.Now()
    tr, err := r.DB.Begin()
    if err != nil {
        return err
    }
    defer tr.Commit()

    // Query the feeding schedule for the given animal
    var scheduleID, lastFedIndex, nextFedIndex int
    err = tr.QueryRow(`SELECT 
        schedule_id, 
        last_fed_index, 
        next_fed_index 
        FROM feeding_schedule 
        WHERE animal_type = $1`, animal).Scan(&scheduleID, &lastFedIndex, &nextFedIndex)
    if err != nil {
        tr.Rollback()
        if err == sql.ErrNoRows {
            return errors.New("feeding schedule not found for the given animal")
        }
        return err
    }

    // Query the schedule times
    var time1, time2, time3 int
    err = tr.QueryRow(`SELECT   
        EXTRACT(HOUR FROM time1) AS hour1, 
        EXTRACT(HOUR FROM time2) AS hour2, 
        EXTRACT(HOUR FROM time3) AS hour3 
        FROM schedules 
        WHERE id = $1`, scheduleID).Scan(&time1, &time2, &time3)
    if err != nil {
        tr.Rollback()
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
        if currentHour >= feedingWindowStart && currentHour <= feedingHour +2 {
            if lastFedIndex == i+1 {
                tr.Rollback()
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

    if (!allowed) {
        nextFeedingWindowStart := (nextFeedingTime - 1 + 24) % 24
        tr.Rollback()
        return fmt.Errorf("next schedule is at %02d:00. You can start feeding animals from %02d:00",
            nextFeedingTime, nextFeedingWindowStart)
    }

    // Check if the provision is suitable for the given animal type
    var provisionAnType string
    var provisionQuantity float64
    err = tr.QueryRow(`SELECT animal_type, quantity FROM provision WHERE type = $1`, provision).Scan(&provisionAnType, &provisionQuantity)
    if err != nil {
        tr.Rollback()
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
        tr.Rollback()
        return errors.New("you cannot feed this provision to that animal")
    }
    // Calculate the total food and water consumption for the animals
    var totalFoodConsumption, totalWaterConsumption float64
    rows, err := tr.Query(`SELECT avg_consumption, avg_water FROM animals WHERE type = $1`, animal)
    if err != nil {
        tr.Rollback()
        return err
    }
    defer rows.Close()


    for rows.Next() {
        var avgConsumption, avgWater float64
        if err := rows.Scan(&avgConsumption, &avgWater); err != nil {
            tr.Rollback()
            return err
        }
        totalFoodConsumption += avgConsumption
        totalWaterConsumption += avgWater
    }

    if err := rows.Err(); err != nil {
        tr.Rollback()
        return err
    }

    // Deduct the food quantity from the provision
    if provisionQuantity < totalFoodConsumption {
        tr.Rollback()
        return errors.New("not enough provision quantity")
    }
    newProvisionQuantity := provisionQuantity - totalFoodConsumption
    _, err = tr.Exec(`UPDATE provision SET quantity = $1 WHERE type = $2`, newProvisionQuantity, provision)
    if err != nil {
        tr.Rollback()
        return err
    }

    // Update the water consumption in the water_consumption table
    _, err = tr.Exec(`UPDATE water_consumption SET total = total + $1`, totalWaterConsumption)
    if err != nil {
        tr.Rollback()
        return err
    }

    // Update the feeding schedule in the database
    _, err = tr.Exec(`UPDATE feeding_schedule SET last_fed_index = $1, next_fed_index = $2 WHERE animal_type = $3`, lastFedIndex, nextFedIndex, animal)
    if err != nil {
        tr.Rollback()
        return err
    }

    return nil
}

func extractWords(str string) []string {
    // Use strings.FieldsFunc for more control over delimiters
    return strings.FieldsFunc(str, func(r rune) bool {
        // Consider characters other than just whitespace for delimiters
        return r == ',' || r == ' ' 
    })
}