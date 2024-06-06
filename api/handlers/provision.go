package handlers

import (
	"farmish/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProvision godoc
// @ID create_provision
// @Router /provision [POST]
// @Summary Create Provision
// @Description Create a new provision
// @Tags Provision
// @Accept json
// @Produce json
// @Param provision body models.BodyProvision true "Provision data"
// @Success 201 {object} string "Provision successfully"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) CreateProvision(c *gin.Context) {
	var body models.BodyProvision
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.PS.CreateProvision(&body); err != nil {
		h.Logger.ERROR.Printf("Failed to create provision: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create provision",
		})
		return
	}

	h.Logger.INFO.Printf("Created provision: %+v", body)
	c.JSON(http.StatusCreated, "Provision successfully created")
}

// GetProvisionById godoc
// @ID get_provision_by_id
// @Router /provision/{id} [GET]
// @Summary Get Provision by ID
// @Description Get a provision by ID
// @Tags Provision
// @Accept json
// @Produce json
// @Param id path string true "Provision ID"
// @Success 200 {object} models.Provision "Provision data"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetProvisionById(c *gin.Context) {
	idPr := c.Param("id")

	var id int
	var err error

	if idPr != "" {
		id, err = strconv.Atoi(idPr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid value for 'id' parameter : %v", err)})
			return
		}
	}

	prs, err := h.Service.PS.GetProvisionById(id)
	if err != nil {
		h.Logger.ERROR.Printf("Failed to get provision: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"provision": prs})
}

// GetAllProvision godoc
// @ID get_all_provisions
// @Router /provision/ [GET]
// @Summary Get All Provisions
// @Description Get all provisions with optional filters
// @Tags Provision
// @Accept json
// @Produce json
// @Param type query string false "Type of provision"
// @Param animal_type query string false "Animal type"
// @Param quantity query string false "Quantity"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} []models.GetProvision "List of provisions"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetAllProvision(c *gin.Context) {
	typ := c.Query("type")
	animal_type := c.Query("animal_type")
	quantityPr := c.Query("quantity")
	limitPr := c.Query("limit")
	offsetPr := c.Query("offset")

	var quantity float64
	var err error
	if quantityPr != "" {
		quantity, err = strconv.ParseFloat(quantityPr, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid value for 'quantity' parameter: %v", err)})
			return
		}
	}

	var limit int
	if limitPr != "" {
		limit, err = strconv.Atoi(limitPr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid value for 'limit' parameter: %v", err)})
			return
		}
	}

	var offset int
	if offsetPr != "" {
		offset, err = strconv.Atoi(offsetPr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid value for 'offset' parameter: %v", err)})
			return
		}
	}

	filter := models.Filter{
		Limit:  limit,
		OFFSET: offset,
	}

	provisions, err := h.Service.PS.GetAllProvision(filter, typ, animal_type, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"provisions": provisions})
}

// UpdateProvision godoc
// @ID update_provision
// @Router /provision/{id} [PUT]
// @Summary Update Provision
// @Description Update a provision by ID
// @Tags Provision
// @Accept json
// @Produce json
// @Param id path string true "Provision ID"
// @Param provision body models.UpdateProvision true "Provision data"
// @Success 200 {object} string "Provision successfully updated"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) UpdateProvision(c *gin.Context) {
	idPr := c.Param("id")
	id, err := strconv.Atoi(idPr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid value for 'id' parameter: %v", err)})
		return
	}

	var body models.UpdateProvision
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.Service.PS.UpdateProvision(&body, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Provision successfully updated!"})
}

// DeleteProvision godoc
// @ID delete_provision
// @Router /provision/{id} [DELETE]
// @Summary Delete Provision
// @Description Delete a provision by ID
// @Tags Provision
// @Accept json
// @Produce json
// @Param id path string true "Provision ID"
// @Success 200 {object} string "Provision successfully deleted"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) DeleteProvision(c *gin.Context) {
	idPr := c.Param("id")

	id, err := strconv.Atoi(idPr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid value for 'id' parameter : %v", err)})
		return
	}

	err = h.Service.PS.DeleteProvision(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Provision successfully deleted!"})
}
