package service

import (
	"farmish/config"
	"farmish/models"
	"farmish/postgresql"
)

type FeedingScheduleService struct {
	FS *postgresql.FeedingScheduleRepo
}

func NewFeedingScheduleService(fs *postgresql.FeedingScheduleRepo) *FeedingScheduleService {
	return &FeedingScheduleService{FS: fs}
}

func (s *FeedingScheduleService) CreateFeedingSchedule(fs *models.FeedingSchedule) (*models.FeedingSchedule, error) {
	ids,err := s.FS.GetAllFeedingScheduleIDs()
	if err != nil {
		return nil,err
	}
	newID,err := config.GenNewID(ids)
	if err!= nil {
        return nil,err
    }
	fs.ID = newID
	return s.FS.CreateFeedingSchedule(fs)
}

func (s *FeedingScheduleService) GetFeedingSchedule(id int) (*models.FeedingSchedule, error) {
    return s.FS.GetFeedingSchedule(id)
}

func (s *FeedingScheduleService) UpdateFeedingSchedule(fs *models.FeedingSchedule) error {
    return s.FS.UpdateFeedingSchedule(fs)
}

func (s *FeedingScheduleService) DeleteFeedingSchedule(id int) error {
    return s.FS.DeleteFeedingSchedule(id)
}