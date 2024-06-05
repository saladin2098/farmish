package service

type Service struct {
	FeedingS FeedingService
	MedS MedicationService
	SheduleS SheduleService
	FeedingScheduleS FeedingScheduleService
}

func NewService(feedingS FeedingService, medS MedicationService, sheduleS SheduleService, feedingScheduleS FeedingScheduleService) *Service {
    return &Service{FeedingS: feedingS, MedS: medS, SheduleS: sheduleS, FeedingScheduleS: feedingScheduleS}
}
