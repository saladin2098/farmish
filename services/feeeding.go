package service

import (
	"farmish/postgresql"
)

type FeedingService struct {
	FR *postgresql.FeedingRepo
}
func NewFeedingService(fr *postgresql.FeedingRepo) *FeedingService {
	return &FeedingService{FR: fr}
}


func (s *FeedingService) FeedAnimals(animal string, provision string) error {
	err :=  s.FR.FeedAnimals(animal, provision)
	if err!= nil {
        return err
    }
	return nil
}