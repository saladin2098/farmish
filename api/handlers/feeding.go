package handlers

import "github.com/gin-gonic/gin"



// FeedAnimals godoc
// @ID feed_animals
// @Summary feeds the animals 
// @Description feeds the given animals given provision 
// @Tags Animal
// @Accept  json
// @Produce  json
// @Param animal query string true "Animal type"
// @Param provision query string true "Provision type"
// @Success 200 {object} string "animals are succesfully fed" 
// @Failure 400 {object} string "Could not feed the animals"
// @Failure 500 {object} string "server error"
// @Router /feeding [GET]
func (h *HTTPHandler) FeedAnimals(c *gin.Context) {
	animal := c.Query("animal")
    provision := c.Query("provision")
    err := h.Service.FeedingS.FeedAnimals(animal, provision)
    if err!= nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not feed the animals:"+err.Error())
		return
    }
    c.JSON(200, gin.H{"status": "animals are succesfully fed"})
}