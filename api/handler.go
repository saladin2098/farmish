package api

import (
	"farmish/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewGin(h *handlers.HTTPHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/feeding", h.FeedAnimals)

	r.POST("/medication", h.CreateMedication)
	r.GET("/medication", h.GetMedication)
	r.PUT("/medication", h.UpdateMedication)
	r.DELETE("/medication/:id", h.DeleteMedication)
	r.GET("/medications", h.GetMedicationsGroupedByType)

	r.POST("/schedule",h.CreateSchedule)
	r.GET("/schedule/:id",h.GetSchedule)
	r.PUT("/schedule",h.UpdateSchedule)
	r.DELETE("/schedule/:id",h.DeleteSchedule)

	r.POST("/feeding_schedule",h.CreateFeedingSchedule)
	r.GET("/feeding_schedule/:id",h.GetFeedingSchedule)
	r.PUT("/feeding_schedule",h.UpdateFeedingSchedule)
	r.DELETE("/feeding_schedule/:id",h.DeleteFeedingSchedule)


	return r
}
