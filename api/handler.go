package api

import (
	"farmish/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewGin(h *handlers.HTTPHandler) *gin.Engine {
	r := gin.Default()

	animal := r.Group("/animal")
	animal.POST("/", h.CreateAnimal)

	return r
}
