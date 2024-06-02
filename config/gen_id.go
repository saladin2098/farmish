package config

import (
	"errors"
	"math/rand"
)

func GenNewID(ids []int) (int, error) {
	if len(ids) >= 90000 {
		return 0, errors.New("number not generated: the list is full")
	}
	for {
		num := rand.Intn(90000) + 10000
		if isNotIn(ids, num) {
			return num, nil
		}
	}
}

func isNotIn(slice []int, id int) bool {
	for _, v := range slice {
		if v == id {
			return false
		}
	}
	return true
}
