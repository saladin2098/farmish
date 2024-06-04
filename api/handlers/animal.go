package handlers

import (
	"database/sql"
	"farmish/models"
	"farmish/postgres/managers"
	"farmish/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	db *sql.DB
)

func (h *HTTPHandler) CreateAnimal(c *gin.Context) {
	var body models.BodyProvision
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	animalService := service.AnimalService{AR: &managers.AnimalRepo{db}}
	createdAnimal, err := animalService.CreateAnimal()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"provision": createdAnimal})
}
