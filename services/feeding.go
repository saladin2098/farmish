package service

import "farmish/postgresql/managers"

type FeedingService struct {
	FR *managers.FeedingRepo
}

func NewFeedingService(fr *managers.FeedingRepo) *FeedingService {
	return &FeedingService{FR: fr}
}

func (s *FeedingService) FeedAnimals(animal string, provision string) error {
	err := s.FR.FeedAnimals(animal, provision)
	if err != nil {
		return err
	}
	return nil
}
