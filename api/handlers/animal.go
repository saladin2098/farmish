package handlers

import (
	"farmish/models"

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
}
