package api

import (
	_ "farmish/api/docs"
	"farmish/api/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewGin(h *handlers.HTTPHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dash := r.Group("/dashboard")
	dash.GET("/get-animals-count", h.GetAnimalsCount)
	dash.GET("/get-avg-weight", h.GetAvgWeight)
	dash.GET("/get-sick-animals", h.GetSickAnimals)
	dash.GET("/get-hungry-animals", h.GetHungryAnimals)
	dash.GET("/check-provision", h.CheckProvision)

	animal := r.Group("/animal")
	animal.POST("/", h.CreateAnimal)
	animal.GET("/:id", h.GetAnimal)
	animal.PUT("/:id", h.UpdateAnimal)
	animal.DELETE("/:id", h.DeleteAnimal)
	r.GET("/animals", h.GetAllAnimals)

	provision := r.Group("/provision")
	provision.POST("/", h.CreateProvision)
	provision.GET("/", h.GetProvisionById)
	provision.GET("/get", h.GetAllProvision)
	provision.PUT("/:id", h.UpdateProvision)
	provision.DELETE("/:id", h.DeleteProvision)

	r.GET("/feeding", h.FeedAnimals)

	r.POST("/medication", h.CreateMedication)
	r.GET("/medication", h.GetMedication)
	r.PUT("/medication", h.UpdateMedication)
	r.DELETE("/medication/:id", h.DeleteMedication)
	r.GET("/medications", h.GetMedicationsGroupedByType)
	r.GET("/medications/all", h.GetAllMedications)

	r.POST("/schedule", h.CreateSchedule)
	r.GET("/schedule/:id", h.GetSchedule)
	r.PUT("/schedule/:id", h.UpdateSchedule)
	r.DELETE("/schedule/:id", h.DeleteSchedule)

	r.POST("/feeding_schedule", h.CreateFeedingSchedule)
	r.GET("/feeding_schedule/:id", h.GetFeedingSchedule)
	r.PUT("/feeding_schedule", h.UpdateFeedingSchedule)
	r.DELETE("/feeding_schedule/:id", h.DeleteFeedingSchedule)

	return r
}
