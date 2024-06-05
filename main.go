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
	fmt.Println(111)
	// fmt.Println(dash.GetAnimalsCount())
	// fmt.Println(dash.GetAvgWeight())
	fmt.Println(dash.GetHungryAnimals())
	// fmt.Println(dash.GetSickAnimals())

	notif := dashboard.NewNotification(*logger, *service)
	notif.SendNotifAboutHungryAnimals()

	fmt.Printf("Server started on port %s\n", config.HTTP_PORT)
	logger.INFO.Println("Server started on port: " + config.HTTP_PORT)
	err = r.Run(config.HTTP_PORT)
	em.CheckErr(err)
}

/*
Azizbek:
	AnimalGetAll (filter)+gin+swagger+testing
	Hayvonlar soni
	O’rtacha vazni(hayvonlar bo’yicha)

Shamsiddin:
	o'zini servicelari + gin
	schedule service + gin
	feed service + gin
	++++++++++ swagger
	++++++++++ testing


+++READMEEE
+++NOTIFICATION (och qolganlar, )
*/
