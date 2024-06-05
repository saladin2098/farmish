package handlers

import "github.com/gin-gonic/gin"


func (h *HTTPHandler) FeedAnimals(c *gin.Context) {
	animal := c.Query("animal")
    provision := c.Query("provision")
    err := h.Service.FeedingS.FeedAnimals(animal, provision)
    if err!= nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
    }
    c.JSON(200, gin.H{"status": "animals are succesfully fed"})
}