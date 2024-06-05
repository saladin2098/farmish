package handlers

import (
	"farmish/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) CreateFeedingSchedule(c *gin.Context) {
	var feedingSchedule models.FeedingSchedule
	err := c.BindJSON(&feedingSchedule)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	res, err := h.Service.FeedingScheduleS.CreateFeedingSchedule(&feedingSchedule)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not create feeding schedule:"+err.Error())
		return
	}
	c.JSON(200, gin.H{"feeding schedule": res})
}

func (h *HTTPHandler) GetFeedingSchedule(c *gin.Context) {
	str_id := c.Param("id")
	id, _ := strconv.Atoi(str_id)
	res, err := h.Service.FeedingScheduleS.GetFeedingSchedule(id)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	c.JSON(200, gin.H{"feeding schedule": res})
}

func (h *HTTPHandler) UpdateFeedingSchedule(c *gin.Context) {
	var feedingSchedule models.FeedingSchedule
    err := c.BindJSON(&feedingSchedule)
    if err!= nil {
        h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
        c.String(400, "Could not bind JSON:"+err.Error())
        return
    }
    err = h.Service.FeedingScheduleS.UpdateFeedingSchedule(&feedingSchedule)
    if err!= nil {
        h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
        c.String(400, "Could not bind JSON:"+err.Error())
        return
    }
    c.JSON(200, gin.H{"status": "feeding schedule is succesfully updated"})
}

func (h *HTTPHandler) DeleteFeedingSchedule(c *gin.Context) {
	str_id := c.Param("id")
    id, _ := strconv.Atoi(str_id)
    err := h.Service.FeedingScheduleS.DeleteFeedingSchedule(id)
    if err!= nil {
        h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
        c.String(400, "Could not bind JSON:"+err.Error())
        return
    }
    c.JSON(200, gin.H{"status": "feeding schedule is succesfully deleted"})
}