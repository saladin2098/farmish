package handlers

import (
	_ "net/http"

	"github.com/gin-gonic/gin"
)

// Get Animals Count godoc
// @ID getAnimalCount
// @Router /dashboard/get-animals-count [GET]
// @Summary Get the count of animals
// @Description Retrieve the total count of animals in the system
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} string "Animals count"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetAnimalsCount(c *gin.Context) {
	a, p, err := h.Dashboard.GetAnimalsCount()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"hayvonlar": a, "parrandlar": p})
}

// Get Average Weight of Animals godoc
// @ID getavgweight
// @Router /dashboard/get-avg-weight [GET]
// @Summary Gets the average weight of animals
// @Description Retrieve the average weight of animals in the system
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} string "Animals count"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetAvgWeight(c *gin.Context) {
	a, p, err := h.Dashboard.GetAvgWeight()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"hayvonlar": a, "parrandlar": p})
}

// func (h *HTTPHandler) GetSickAnimals(c *gin.Context) {

// }
