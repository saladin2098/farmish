package dashboard

import (
	"farmish/config/logger"
	"farmish/models"
	service "farmish/services"
	"time"

	"github.com/go-co-op/gocron"
)

type Notification struct {
	Logger    logger.Logger
	Scheduler *gocron.Scheduler
	Service   service.Service
}

func NewNotification(logger logger.Logger, service service.Service) *Notification {
	return &Notification{Logger: logger, Scheduler: gocron.NewScheduler(time.UTC), Service: service}
}

func (n *Notification) SendNotifAboutHungryAnimals() {
	dash := NewDashboard(n.Service)
	animals, err := dash.GetHungryAnimals()
	if err != nil {
		return
	}
	if animals.Count > 0 {
		n.Scheduler.Every(1).Seconds().Do(n.warnHungryAnimals, animals)
		n.Scheduler.StartAsync()
		select {}
	}
}

func (n *Notification) warnHungryAnimals(animals *models.AnimalsGetAll) {
	n.Logger.WARN.Printf("%d ta hayvonlar och qoldi", animals.Count)
}

func (n Notification) warnSickAnimals(animals *models.AnimalsGetAll) {
	n.Logger.WARN.Printf("")
}
