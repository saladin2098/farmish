package service

import "farmish/postgresql/managers"

type HealthConditionService struct {
	HR *managers.HealthConditionRepo
}

func NewHealthConditionService(hr *managers.HealthConditionRepo) *HealthConditionService {
	return &HealthConditionService{HR: hr}
}

func (s *HealthConditionService) GetAllHealthConditionIds() ([]int, error) {
	return s.HR.GetAllHealthConditionIds()
}