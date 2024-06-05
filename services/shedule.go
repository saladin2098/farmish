package service

import (
	"farmish/config"
	"farmish/models"
	"farmish/postgresql"
)

type SheduleService struct {
	SR *postgresql.ScheduleRepo
}

func NewSheduleService(sr *postgresql.ScheduleRepo) *SheduleService {
    return &SheduleService{SR: sr}
}

func (s *SheduleService) CreateSchedule(schedule *models.Schedule) (*models.Schedule,error) {
	SchedileIDs,err := s.SR.GetAllScheduleIDs()
	if err!= nil {
		return nil,err
	}
	newID1,err := config.GenNewID(*SchedileIDs)
	if err!= nil {
		return nil,err
	}
	schedule.ID = newID1
	res,err := s.SR.CreateSchedule(schedule)
	if err!= nil {
        return nil,err
    }
	return res,nil
}
func (s *SheduleService) GetShedule(id int) (*models.Schedule, error) {
	res,err := s.SR.GetSchedule(id)
	if err!= nil {
		return nil,err
	}
	return res,nil
}
func (s *SheduleService) UpdateSchedule(shedule *models.Schedule) error {
	err := s.SR.UpdateSchedule(shedule)
	if err != nil {
		return err
	}
	return nil
}
func (s *SheduleService) DeleteSchedule(id int) error {
	err := s.SR.DeleteSchedule(id)
    if err!= nil {
        return err
    }
    return nil
}