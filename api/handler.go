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

	animal := r.Group("/animal")
	animal.POST("/", h.CreateAnimal)
	r.GET("/animals", h.GetAllAnimals)

	provision := r.Group("/provision")
	provision.POST("/", h.CreateProvision)
	provision.GET("/", h.GetProvision)
	provision.GET("/get", h.GetAllProviison)
	provision.PUT("/:id", h.UpdateProvision)
	provision.DELETE("/:id", h.DeleteProvision)

	return r
}
