package main

import (
	cf "farmish/config"
	"farmish/config/logger"
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
	defer db.Close()
	repo := postgresql.NewMidacationRepo(db)
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

	med, err := repo.GetMedicationsGroupedByType()
	if err!= nil {
        em.CheckErr(err)
    }

	fmt.Println(med)
}
