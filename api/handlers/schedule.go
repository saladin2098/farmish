package handlers

import (
	"farmish/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateSchedule godoc
// @ID create_schedule
// @Summary Create a new schedule
// @Description Create a new schedule
// @Tags Schedule
// @Accept json
// @Produce json
// @Param schedule body models.ScheduleCreate true "Schedule data"
// @Success 200 {object} models.Schedule
// @Failure 400 string string
// @Failure 500 {object} models.Schedule
// @Router /schedule [POST]
func (h *HTTPHandler) CreateSchedule(c *gin.Context) {
	var sc models.ScheduleCreate
	err := c.BindJSON(&sc)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	timeLayout := "15:04:05"
	time1, _ := time.Parse(timeLayout, sc.Time1)
	time2, _ := time.Parse(timeLayout, sc.Time2)
	time3, _ := time.Parse(timeLayout, sc.Time3)

	var schedule models.Schedule
	schedule.Time1 = time1
	schedule.Time2 = time2
	schedule.Time3 = time3

	res, err := h.Service.SheduleS.CreateSchedule(&schedule)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	c.JSON(200, gin.H{"schedule": res})
}

// GetSchedule godoc
// @Summary Get a schedule by ID
// @Description Get a schedule by ID
// @Tags Schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 200 {object}  models.Schedule
// @Failure 400 {string} string "Could not bind JSON"
// @Router /schedule/{id} [GET]
func (h *HTTPHandler) GetSchedule(c *gin.Context) {
	str_id := c.Param("id")
	id, _ := strconv.Atoi(str_id)
	res, err := h.Service.SheduleS.GetShedule(id)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	c.JSON(200, gin.H{"schedule": res})
}

// UpdateSchedule godoc
// @Summary Update a schedule
// @Description Update a schedule
// @Tags Schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Param schedule body models.ScheduleCreate true "Schedule data"
// @Success 200 {string} string "Schedule is successfully updated"
// @Failure 400 {string} string "Could not bind JSON"
// @Router /schedule/{id} [PUT]
func (h *HTTPHandler) UpdateSchedule(c *gin.Context) {
	var sc models.ScheduleCreate
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	err := c.BindJSON(&sc)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	timeLayout := "15:04:05"
	time1, _ := time.Parse(timeLayout, sc.Time1)
	time2, _ := time.Parse(timeLayout, sc.Time2)
	time3, _ := time.Parse(timeLayout, sc.Time3)

	var schedule models.Schedule
	schedule.ID = id
	schedule.Time1 = time1
	schedule.Time2 = time2
	schedule.Time3 = time3

	err = h.Service.SheduleS.UpdateSchedule(&schedule)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "schedule is succesfully updated"})
}

// DeleteSchedule godoc
// @Summary Delete a schedule by ID
// @Description Delete a schedule by ID
// @Tags Schedule
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 200 {string} string "Schedule is successfully deleted"
// @Failure 400 {string} string "Could not bind JSON"
// @Router /schedule/{id} [DELETE]
func (h *HTTPHandler) DeleteSchedule(c *gin.Context) {
	str_id := c.Param("id")
	id, _ := strconv.Atoi(str_id)
	err := h.Service.SheduleS.DeleteSchedule(id)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	c.JSON(200, gin.H{"status": "schedule is succesfully deleted"})
}
