package handlers

import (
	"farmish/models"
	"strconv"

	"github.com/gin-gonic/gin"
)
// CreateFeedingSchedule godoc
// @ID create_feeding_schedule
// @Router /feeding_schedule [POST]
// @Summary Create Feeding Schedule
// @Description Create a new feeding schedule
// @Tags FeedingSchedule
// @Accept json
// @Produce json
// @Param feedingSchedule body models.FeedingSchedule true "Feeding Schedule data"
// @Success 200 {object} models.FeedingSchedule
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
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
		h.Logger.ERROR.Printf("Could not create f schedule: %s", err.Error())
		c.String(400, "Could not create feeding schedule:"+err.Error())
		return
	}
	c.JSON(200, gin.H{"feeding schedule": res})
}
// GetFeedingSchedule godoc
// @ID get_feeding_schedule
// @Router /feeding_schedule/{id} [GET]
// @Summary Get Feeding Schedule by ID
// @Description Get a feeding schedule by ID
// @Tags FeedingSchedule
// @Accept json
// @Produce json
// @Param id path string true "Feeding Schedule ID"
// @Success 200 {object}  models.FeedingSchedule
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
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

// UpdateFeedingSchedule godoc
// @ID update_feeding_schedule
// @Router /feeding_schedule [PUT]
// @Summary Update Feeding Schedule
// @Description Update a feeding schedule by ID
// @Tags FeedingSchedule
// @Accept json
// @Produce json
// @Param feedingSchedule body models.FeedingSchedule true "Feeding Schedule data"
// @Success 200 {object} string "feeding schedule is successfully updated"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
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
// DeleteFeedingSchedule godoc
// @ID delete_feeding_schedule
// @Router /feeding_schedule/{id} [DELETE]
// @Summary Delete Feeding Schedule
// @Description Delete a feeding schedule by ID
// @Tags FeedingSchedule
// @Accept json
// @Produce json
// @Param id path string true "Feeding Schedule ID"
// @Success 200 {object} string "feeding schedule is successfully deleted"
// @Response 400 {object} string "Bad Request"
// @Failure 500 {object} string "Server Error"
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