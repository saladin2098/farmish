package handlers

import (
	"farmish/models"
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
		c.JSON(500, gin.H{"error while getting weights: ": err.Error()})
		return
	}
	c.JSON(200, gin.H{"hayvonlar": a, "parrandlar": p})
}

// Get Sick Animals godoc
// @ID getsickanimals
// @Router /dashboard/get-sick-animals [GET]
// @Summary Gets the count and list of sick animals
// @Description Retrieve the list of sick animals
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} string "Animals:"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetSickAnimals(c *gin.Context) {
	animals, err := h.Dashboard.GetSickAnimals()
	if err != nil {
		c.JSON(500, gin.H{"error while getting sick animals: ": err.Error()})
		return
	}
	c.JSON(200, animals)
}

// Get Hungry Animals godoc
// @ID gethungryanimals
// @Router /dashboard/get-hungry-animals [GET]
// @Summary Gets the count and list of hungry animals
// @Description Retrieve the list of hungry animals
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} string "Animals:"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetHungryAnimals(c *gin.Context) {
	animals, err := h.Dashboard.GetHungryAnimals()
	if err != nil {
		c.JSON(500, gin.H{"error while getting hungry animals: ": err.Error()})
		return
	}
	c.JSON(200, animals)
}

// Check provision godoc
// @ID checkprovision
// @Router /dashboard/check-provision [GET]
// @Summary Checks quantity of provision and returns if it's enough for 6 days
// @Description Retrieve provision data
// @Tags Dashboard
// @Accept json
// @Produce json
// @Success 200 {object} string "Animals:"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) CheckProvision(c *gin.Context) {
	animals, poultries, err := h.Dashboard.CheckProvision()
	if err != nil {
		c.JSON(500, gin.H{"error while checking provision: ": err.Error()})
		return
	}
	pr_animals, err := h.Service.PS.PR.GetAllProvision(models.Filter{Limit: 100000, OFFSET: 0}, "", "hayvon", 0)
	if err != nil {
		c.JSON(500, gin.H{"error getting provision: ": err.Error()})
		return
	}
	pr_poultries, err := h.Service.PS.PR.GetAllProvision(models.Filter{Limit: 10000, OFFSET: 0}, "", "parranda", 0)
	if err != nil {
		c.JSON(500, gin.H{"error getting provision: ": err.Error()})
		return
	}
	res := map[string]interface{}{
		"hayvonlar":                       pr_animals[0].Quantity,
		"hayvonlar yemishi yetarliligi":   animals,
		"parrandalar":                     pr_poultries[0].Quantity,
		"parrandalar yemishi yetarliligi": poultries,
	}
	c.JSON(200, res)
}
