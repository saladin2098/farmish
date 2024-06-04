package postgresql

import (
	"database/sql"
	// "time"
)

type FeedingRepo struct {
	DB *sql.DB
}

func NewFeedingRepo(db *sql.DB) *FeedingRepo {
	return &FeedingRepo{DB: db}
}

// func (r *FeedingRepo) FeedAnimals(animal string) error {
// 	current_time := time.Now().Hour()
// }
