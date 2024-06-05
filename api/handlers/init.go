package handlers

import (
	"farmish/config/logger"
	"farmish/dashboard"
	service "farmish/services"
)

type HTTPHandler struct {
	Service   *service.Service
	Dashboard *dashboard.Dashboard
	Logger    logger.Logger
}

func NewHTTPHandler(service *service.Service, dashboard *dashboard.Dashboard, logger logger.Logger) *HTTPHandler {
	return &HTTPHandler{Service: service, Dashboard: dashboard, Logger: logger}
}
