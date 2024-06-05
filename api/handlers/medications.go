package handlers

import (
	"farmish/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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