package service

type Service struct {
	AR               AnimalService
	PS               ProvisionService
	HR               HealthConditionService
	FeedingS         FeedingService
	MedS             MedicationService
	SheduleS         SheduleService
	FeedingScheduleS FeedingScheduleService
}

func NewService(
	as AnimalService, hs HealthConditionService, ps ProvisionService,
	feedingS FeedingService, medS MedicationService, sheduleS SheduleService, feedingScheduleS FeedingScheduleService,
) *Service {
	return &Service{
		AR: as, HR: hs, PS: ps,
		FeedingS: feedingS, MedS: medS, SheduleS: sheduleS, FeedingScheduleS: feedingScheduleS,
	}
}
