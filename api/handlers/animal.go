package handlers

import (
	"farmish/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) CreateAnimal(c *gin.Context) {
	var animal models.AnimalCreate
	if err := c.ShouldBindJSON(&animal); err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	if err := h.Service.AR.CreateAnimal(&animal); err != nil {
		h.Logger.ERROR.Printf("Could not create animal: %s", err.Error())
		c.String(500, "Could not create animal:"+err.Error())
		return
	}
	h.Logger.INFO.Printf("Animal created: %+v", animal)
	c.String(200, "Animal created")
}

func (h *HTTPHandler) GetAnimal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.ERROR.Printf("Could not parse id: %s", err.Error())
		c.String(400, "Could not parse id:"+err.Error())
		return
	}
	animal, err := h.Service.AR.GetAnimal(id)
	if err != nil {
		h.Logger.ERROR.Printf("Could not get animal: %s", err.Error())
		c.String(500, "Could not get animal:"+err.Error())
		return
	}
	c.JSON(200, animal)
}

func (h *HTTPHandler) GetAllAnimals(c *gin.Context) {
	animal_type := c.Query("type")
	is_healthy := c.Query("is_healthy")
	is_hungry := c.Query("is_hungry")

	animals, err := h.Service.AR.GetAllAnimals(animal_type, is_healthy, is_hungry)
	if err != nil {
		h.Logger.ERROR.Printf("Could not get all animals: %s", err.Error())
		c.String(500, "Could not get all animals:"+err.Error())
		return
	}
	c.JSON(200, animals)
}

func (h *HTTPHandler) UpdateAnimal(c *gin.Context) {
	animal := models.AnimalUpdate{}

	if err := c.BindJSON(&animal); err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.ERROR.Printf("Could not parse id: %s", err.Error())
		c.String(400, "Could not parse id:"+err.Error())
		return
	}
	animal.ID = int32(id)

	if err := h.Service.AR.UpdateAnimal(&animal); err != nil {
		h.Logger.ERROR.Printf("Could not update animal: %s", err.Error())
		c.String(500, "Could not update animal:"+err.Error())
		return
	}
	h.Logger.INFO.Printf("Animal updated: %+v", animal)
	c.String(200, "Animal updated")
}

func (h *HTTPHandler) DeleteAnimal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.ERROR.Println("Could not parse id" + err.Error())
		c.String(400, "Could not parse id:"+err.Error())
		return
	}
	if err := h.Service.AR.DeleteAnimal(id); err != nil {
		h.Logger.ERROR.Println("Could not delete animal" + err.Error())
		c.String(500, "Could not delete animal:"+err.Error())
		return
	}
	h.Logger.INFO.Printf("Animal deleted: %d", id)
	c.String(200, "Animal deleted")
}
