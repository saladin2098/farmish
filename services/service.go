package service

type Service struct {
	AR AnimalService
	HR HealthConditionService
}

func NewService(ar AnimalService, hr HealthConditionService) *Service {
	return &Service{AR: ar, HR: hr}
}
