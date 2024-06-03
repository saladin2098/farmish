package main

import (
	cf "farmish/config"
	"farmish/config/logger"
	"farmish/postgresql"
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

	// ar := managers.NewAnimalRepo(db)
	// ids, err := ar.GetAllAnimalIds()
	// if err != nil {
	// 	logger.ERROR.Panicln(err)
	// }
	// newid, err := cf.GenNewID(ids)
	// if err != nil {
	// 	logger.ERROR.Panicln(err)
	// }

	// ar.UpdateAnimal(&models.AnimalUpdate{
	// 	ID: ,
	// })
	// sr := managers.NewScheduleRepo(db)

	// ids, err := sr.GetAllScheduleIds()
	// if err != nil {
	// 	logger.ERROR.Panicln(err)
	// }
	// newid, err := cf.GenNewID(ids)
	// if err != nil {
	// 	logger.ERROR.Panicln(err)
	// }
	// err = sr.CreateSchedule(&models.Schedule{
	// 	ID:         newid,
	// 	AnimalType: "o'rdak",
	// 	Time1:      time.Time{}.Add(time.Hour * 7),
	// 	Time2:      time.Time{}.Add(time.Hour * 13),
	// 	Time3:      time.Time{}.Add(time.Hour * 19),
	// })
	// if err != nil {
	// 	logger.ERROR.Panicln(err)
	// }

	// schedule := models.Schedule{
	// 	ID:         1,
	// 	AnimalType: "qush",
	// 	Time1:      time.Time{}.Add(time.Hour * 7),
	// 	Time2:      time.Time{}.Add(time.Hour * 13),
	// 	Time3:      time.Time{}.Add(time.Hour * 19),
	// }
	// feedSchedule := models.FeedingSchedule{
	// 	ID:           1,
	// 	LastFedIndex: 1,
	// }
	// sr.CreateSchedule(&schedule, &feedSchedule)
}
