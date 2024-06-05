package main

import (
	"farmish/api"
	"farmish/api/handlers"
	cf "farmish/config"
	"farmish/config/logger"
	service "farmish/services"
	// m "farmish/models"
	"farmish/postgresql"
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

	// defer db.Close()
	// repo := postgresql.NewMidacationRepo(db)
	// med, err := repo.CreateMedication(&m.Medications{
	// 	ID:        4,
	//     Name:      "birnima4",
	//     Type:      "qorasonga",
	//     Quantity:  10,
	// })
	// if err!= nil {
	//     em.CheckErr(err)
	// }

	// med,err := repo.GetMedication(0,"","tablet")
	// if err != nil {
	// 	em.CheckErr(err)
	// }

	// med, err := repo.GetMedicationsGroupedByType("tablet")
	// if err != nil {
	// 	em.CheckErr(err)
	// }
	// time := time.Now().Hour

	// fmt.Println(time)

	// schedule := postgresql.NewFeedingScheduleRepo(db)
	// f := m.FeedingSchedule{
	// 	ID:        2,
    //     AnimalType: "tuya",
    //     LastFedIndex: 2,
    //     NextFedIndex: 3,
    //     ScheduleID: 63231,
	// }
	// fs,err := schedule.CreateFeedingSchedule(&f)
	// if err!= nil {
    //     em.CheckErr(err)
    // }
	// fmt.Println(fs)

	// feedingREpo := postgresql.NewFeedingRepo(db)
	// err = feedingREpo.FeedAnimals("ot","hay")
	// if err!= nil {
    //     em.CheckErr(err)
    // }

	// fmt.Println(med)
		
	fs := service.NewFeedingService(postgresql.NewFeedingRepo(db))
	fss := service.NewFeedingScheduleService(postgresql.NewFeedingScheduleRepo(db))
	ms := service.NewMedicationService(postgresql.NewMidacationRepo(db))
    ss := service.NewSheduleService(postgresql.NewScheduleRepo(db))

	service := service.NewService(*fs,*ms,*ss,*fss)
	h := handlers.NewHTTPHandler(service, *logger)
	r := api.NewGin(h)

	fmt.Printf("Server started on port %s\n", config.HTTP_PORT)
	logger.INFO.Println("Server started on port: " + config.HTTP_PORT)
	err = r.Run(config.HTTP_PORT)
	em.CheckErr(err)




}
