package main

import (
	"farmish/api"
	"farmish/api/handlers"
	"farmish/config"
	"farmish/config/logger"
	p "farmish/postgres"
	"farmish/service"
	"strconv"
)

func main() {
	cfg := config.Load()

	log := logger.NewLogger("logs", "log.txt")

	db, err := p.ConnectDB(cfg)
	if err != nil {
		log.ERROR.Fatalln("Error connecting to the database:", err)
	}

	srv := service.NewService(db)
	h := handlers.NewHTTPHandler(srv, *log)
	r := api.NewGin(h)

	r.Run(":" + strconv.Itoa(cfg.HTTP_PORT))
}
