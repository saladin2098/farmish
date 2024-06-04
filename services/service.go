package service

type Service struct {
	AR AnimalService
	PS ProvisionService
	HR HealthConditionService
}

func NewService(as AnimalService, hs HealthConditionService, ps ProvisionService) *Service {
	return &Service{AR: as, HR: hs, PS: ps}
}
