package handlers

import (
	"farmish/models"
	_ "net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAnimal godoc
// @ID create_animal
// @Router /animal [POST]
// @Summary Create Animal
// @Description Create a new animal
// @Tags Animal
// @Accept json
// @Produce json
// @Param animal body models.AnimalCreate true "Animal data"
// @Success 200 {object} string "Animal created"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
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

// GetAnimal godoc
// @ID get_animal
// @Router /animal/{id} [GET]
// @Summary Get Animal
// @Description Get an animal by ID
// @Tags Animal
// @Accept json
// @Produce json
// @Param id path string true "Animal ID"
// @Success 200 {object} models.AnimalGet "Animal data"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
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

// GetAllAnimals godoc
// @ID get_all_animals
// @Router /animals [GET]
// @Summary Get All Animals
// @Description Get all animals with optional filters
// @Tags Animal
// @Accept json
// @Produce json
// @Param type query string false "Animal type"
// @Param is_healthy query string false "Is Healthy"
// @Param is_hungry query string false "Is Hungry"
// @Success 200 {object} models.AnimalsGetAll "Animals data"
// @Failure 500 {object} string "Server Error"
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

// UpdateAnimal godoc
// @ID update_animal
// @Router /animal/{id} [PUT]
// @Summary Update Animal
// @Description Update an animal's information
// @Tags Animal
// @Accept json
// @Produce json
// @Param id path string true "Animal ID"
// @Param animal body models.AnimalUpdate true "Animal data"
// @Success 200 {object} string "Animal updated"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
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

// DeleteAnimal godoc
// @ID delete_animal
// @Router /animal/{id} [DELETE]
// @Summary Delete Animal
// @Description Delete an animal by ID
// @Tags Animal
// @Accept json
// @Produce json
// @Param id path string true "Animal ID"
// @Success 200 {object} string "Animal Deleted"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) DeleteAnimal(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.ERROR.Printf("Could not parse id: %s", err.Error())
		c.String(400, "Could not parse id:"+err.Error())
		return
	}
	if err := h.Service.AR.DeleteAnimal(id); err != nil {
		h.Logger.ERROR.Printf("Could not update animal: %s", err.Error())
		c.String(500, "Could not update animal:"+err.Error())
		return
	}
	h.Logger.INFO.Printf("Animal deleted: %+v", id)
	c.String(200, "Animal deleted")
}
