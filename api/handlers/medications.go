package handlers

import (
	"farmish/models"
	"strconv"

	"github.com/gin-gonic/gin"
)
// CreateMedication godoc
// @Summary Create a new medication
// @Description Create a new medication
// @Tags Medication
// @Accept json
// @Produce json
// @Param medication body models.MedicationsGet true "Medication data"
// @Success 200 {object} models.Medications
// @Failure 400 {string} string "Could not bind JSON"
// @Router /medication [POST]
func (h *HTTPHandler) CreateMedication(c *gin.Context) {
	var medication models.MedicationsGet
	err := c.BindJSON(&medication)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	med, err := h.Service.MedS.CreateMedication(&medication)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	c.JSON(200, med)
}
// GetMedication godoc
// @Summary Get a medication by ID, name, or type
// @Description Get a medication by ID, name, or type
// @Tags Medication
// @Accept json
// @Produce json
// @Param id query int false "Medication ID"
// @Param name query string false "Medication Name"
// @Param type query string false "Medication Type"
// @Success 200 {object} models.Medications
// @Failure 400 {string} string "Could not bind JSON"
// @Router /medication [GET]
func (h *HTTPHandler) GetMedication(c *gin.Context) {
	str_id := c.Query("id")
	name := c.Query("name")
	turi := c.Query("type")
	id, _ := strconv.Atoi(str_id)
	med, err := h.Service.MedS.GetMedication(id, name, turi)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	c.JSON(200, med)
}
// DeleteMedication godoc
// @Summary Delete a medication by ID
// @Description Delete a medication by ID
// @Tags Medication
// @Accept json
// @Produce json
// @Param id path int true "Medication ID"
// @Success 200 {string} string "Medication is successfully deleted"
// @Failure 400 {string} string "Could not bind JSON"
// @Router /medication/{id} [DELETE]
func (h *HTTPHandler) DeleteMedication(c *gin.Context) {
	str_id := c.Param("id")
    id, _ := strconv.Atoi(str_id)
    err := h.Service.MedS.DeleteMedication(id)
    if err!= nil {
        h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
        c.String(400, "Could not bind JSON:"+err.Error())
        return
    }
    c.JSON(200, gin.H{"status": "medication is succesfully deleted"})
}
// UpdateMedication godoc
// @Summary Update a medication
// @Description Update a medication
// @Tags Medication
// @Accept json
// @Produce json
// @Param medication body models.Medications true "Medication data"
// @Success 200 {string} string "medication is successfully updated"
// @Failure 400 {object} string "Could not bind JSON"
// @Failure 500 {object} string  "Server error"
// @Router /medication [PUT]
func (h *HTTPHandler) UpdateMedication(c *gin.Context) {
	var medication models.Medications
    err := c.BindJSON(&medication)
    if err!= nil {
        h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
        c.String(400, "Could not bind JSON:"+err.Error())
        return
    }
    err = h.Service.MedS.UpdateMedication(&medication)
    if err!= nil {
        h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
        c.String(400, "Could not bind JSON:"+err.Error())
        return
    }
    c.JSON(200, gin.H{"status": "medication is succesfully updated"})
}
// GetMedicationsGroupedByType godoc
// @Summary Get medications grouped by type
// @Description Get medications grouped by type
// @Tags Medication
// @Accept json
// @Produce json
// @Param type query string true "Medication Type"
// @Success 200 {object} models.MedicarionaGrouped
// @Failure 400 {string} string "Could not bind JSON"
// @Router /medications [GET]
func (h *HTTPHandler) GetMedicationsGroupedByType(c *gin.Context) {
	tur := c.Query("type")
    med, err := h.Service.MedS.GetMedicationsGroupedByType(tur)
    if err!= nil {
        h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
        c.String(400, "Could not bind JSON:"+err.Error())
        return
    }
    c.JSON(200, med)
}