package main

import (
	"farmish/api"
	"farmish/api/handlers"
	cf "farmish/config"
	"farmish/config/logger"
	"farmish/dashboard"
	"farmish/postgresql"
	"farmish/postgresql/managers"
	service "farmish/services"
	"fmt"
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH)
	em := cf.NewErrorManager(logger)

	db, err := postgresql.ConnectDB(config)
	em.CheckErr(err)
	defer db.Close()

	ar := managers.NewAnimalRepo(db)
	hr := managers.NewHealthConditionRepo(db)
	pr := managers.NewProvisionRepo(db)

	as := service.NewAnimalService(hr, ar)
	hs := service.NewHealthConditionService(hr)
	ps := service.NewProvisionService(pr)
	service := service.NewService(*as, *hs, *ps)
	h := handlers.NewHTTPHandler(service, *logger)
	r := api.NewGin(h)

	dash := dashboard.NewDashboard(*service)
	notif := dashboard.NewNotification(*logger, *dash)

	fmt.Println(1111)
	go notif.SendNotifAboutHungryAnimals()
	go notif.SendNotifAboutSickAnimals()

	fmt.Printf("Server started on port %s\n", config.HTTP_PORT)
	logger.INFO.Println("Server started on port: " + config.HTTP_PORT)
	err = r.Run(config.HTTP_PORT)
	em.CheckErr(err)
}
