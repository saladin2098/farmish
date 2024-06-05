package handlers

import (
	"farmish/config/logger"
	service "farmish/services"
)

type HTTPHandler struct {
	Service *service.Service
	Logger  logger.Logger
}

func NewHTTPHandler(service *service.Service, logger logger.Logger) *HTTPHandler {
	return &HTTPHandler{Service: service, Logger: logger}
}
