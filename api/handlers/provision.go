package handlers

import (
	"database/sql"
	"farmish/models"
	"farmish/postgresql/managers"
	service "farmish/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	db *sql.DB
)

// CreateProvision godoc
// @ID create_provision
// @Router /provision [POST]
// @Summary Create Provision
// @Description Create Provision
// @Tags Provision
// @Accept json
// @Produce json
// @Param user body models.BodyProvision true "Created Provision"
// @Success 201 {object} string "Provision data"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) CreateProvision(c *gin.Context) {
	var body models.BodyProvision
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var provisionService = service.ProvisionService{PR: &managers.ProvisionRepo{}}
	createdProvision, err := provisionService.CreateProvision(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"provision": createdProvision})
}

// GetProvisionById godoc
// @ID get_provision_by_id
// @Router /rovision/{id}{type}{animal_type}{quantity} [GET]
// @Summary Get Provision By ID
// @Description Get a provision by ID
// @Tags Provision
// @Accept json
// @Produce json
// @Param id path string true "Provision ID"
// @Success 200 {object} models.GetProvision "Provision data"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetProvision(c *gin.Context) {
	idPr := c.Param("id")
	typ := c.Param("type")
	animal_type := c.Param("animal_type")
	quantityPr := c.Param("quantity")

	id, err := strconv.Atoi(idPr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid value for 'id' parameter : %v", err)})
		return
	}

	quantity, err := strconv.ParseFloat(quantityPr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid value for 'quantity' parameter : %v", err)})
		return
	}

	prs, err := h.Service.PS.GetProvision(id, typ, animal_type, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"provision": prs})
}

// GetAllProvisions godoc
// @ID get_all_provisions
// @Router /provision [GET]
// @Summary Get All Provisions
// @Description Get All Provisions
// @Tags Provision
// @Accept json
// @Produce json
// @Success 200 {object} models.GetAllProvisions "Provisions data"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) GetAllProviison(c *gin.Context) {
	var body models.Filter
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	provisions, err := h.Service.PS.GetAllProvision(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"provisions": provisions})
}

// UpdateProvision godoc
// @ID update_provision
// @Router /provision/ [PUT]
// @Summary Update Provision
// @Description Update a provision
// @Tags Provision
// @Accept json
// @Produce json
// @Param id path string true "Provision ID"
// @Param provision body models.UpdateProvision true "Provision data"
// @Success 200 {object} string "Provision updated successfully"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
func (h *HTTPHandler) UpdateProvision(c *gin.Context) {
	var body models.UpdateProvision
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	provisions, err := h.Service.PS.UpdateProvision(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"provisions": provisions})
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
// @Success 200 {object} string "Provision deleted successfully"
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

	c.JSON(http.StatusNoContent, "")
}
