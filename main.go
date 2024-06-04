package main

import (
	"farmish/models"
	"farmish/service"
	"fmt"
)

func main() {
	//cfg := config.Load()
	//
	//log := logger.NewLogger("logs", "log.txt")
	//
	//db, err := p.ConnectDB(cfg)
	//if err != nil {
	//	log.ERROR.Fatalln("Error connecting to the database:", err)
	//}
	//
	//srv := service.NewService(db)
	//h := handlers.NewHTTPHandler(srv, *log)
	//r := api.NewGin(h)
	//
	//r.Run()

	pr := models.BodyProvision{
		AnimalType: "Mammals",
		Type:       "Hay",
		Quantity:   10,
	}

	s := service.Service{}
	prs, err := s.CreateProvision(&pr)

	fmt.Println(prs, err)
}
