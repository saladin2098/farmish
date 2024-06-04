package service

type Service struct {
	AR AnimalService
	HR HealthConditionService
}

func NewService(as AnimalService, hs HealthConditionService) *Service {
	return &Service{AR: as, HR: hs}
}
