package handlers

import (
	"farmish/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) CreateSchedule(c *gin.Context) {
	var sc models.ScheduleCreate
	err := c.BindJSON(&sc)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	timeLayout := "15:04:05"
	time1,_ := time.Parse(timeLayout, sc.Time1)
	time2,_ := time.Parse(timeLayout, sc.Time2)
	time3,_ := time.Parse(timeLayout, sc.Time3)

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
func (h *HTTPHandler) UpdateSchedule(c *gin.Context) {
	var sc models.ScheduleCreate
	err := c.BindJSON(&sc)
	if err != nil {
		h.Logger.ERROR.Printf("Could not bind JSON: %s", err.Error())
		c.String(400, "Could not bind JSON:"+err.Error())
		return
	}
	timeLayout := "15:04:05"
	time1,_ := time.Parse(timeLayout, sc.Time1)
	time2,_ := time.Parse(timeLayout, sc.Time2)
	time3,_ := time.Parse(timeLayout, sc.Time3)

	var schedule models.Schedule
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
