package dashboard

import (
	"farmish/config/logger"
	"time"

	"github.com/go-co-op/gocron"
)

type Notification struct {
	Logger    logger.Logger
	Scheduler *gocron.Scheduler
	Dashboard Dashboard
}

func NewNotification(logger logger.Logger, dashboard Dashboard) *Notification {
	return &Notification{Logger: logger, Scheduler: gocron.NewScheduler(time.UTC), Dashboard: dashboard}
}

func (n *Notification) SendNotifAboutHungryAnimals() {
	n.Scheduler.Every(5).Seconds().Do(func() {
		err := n.warnHungryAnimals()
		if err != nil {
			n.Logger.ERROR.Println(err)
		}
	})
	n.Scheduler.StartAsync()
}

func (n *Notification) SendNotifAboutSickAnimals() {
	n.Scheduler.Every(5).Seconds().Do(func() {
		err := n.warnSickAnimals()
		if err != nil {
			n.Logger.ERROR.Println(err)
		}
	})
	n.Scheduler.StartAsync()
}

func (n *Notification) warnHungryAnimals() error {
	animals, err := n.Dashboard.GetHungryAnimals()
	if err != nil {
		return err
	}
	if animals.Count > 0 {
		n.Logger.WARN.Printf("%d ta hayvonlar och qoldi!!!", animals.Count)
	}
	return nil
}

func (n *Notification) warnSickAnimals() error {
	animals, err := n.Dashboard.GetSickAnimals()
	if err != nil {
		return err
	}
	if animals.Count > 0 {
		n.Logger.WARN.Printf("%d ta hayvonlar kasal!!!", animals.Count)
	}
	return nil
}

func (n *Notification) SendNotifAboutProvision() {
	n.Scheduler.Every(5).Seconds().Do(func() {
		err := n.warnProvision()
		if err != nil {
			n.Logger.ERROR.Println(err)
		}
	})
	n.Scheduler.StartAsync()
}

func (n *Notification) warnProvision() error {
	animals, poultry, err := n.Dashboard.CheckProvision()
	if err != nil {
		return err
	}
	if !animals {
		n.Logger.WARN.Println("Hayvonlarning yemishi oz qoldi")
	}
	if !poultry {
		n.Logger.WARN.Println("Parrandalarning yemishi oz qoldi")
	}
	return nil
}
