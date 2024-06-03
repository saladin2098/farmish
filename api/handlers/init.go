package handlers

import (
	"farmish/config/logger"
	"farmish/service"
)

type HTTPHandler struct {
	Logger  logger.Logger
	Service *service.Service
}

func NewHTTPHandler(service *service.Service, logger logger.Logger) *HTTPHandler {
	return &HTTPHandler{Service: service, Logger: logger}

}
