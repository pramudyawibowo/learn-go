package controller

import (
	"fmt"
	"sesi7/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type WaterWindController struct {
	db *gorm.DB
}

func NewWaterWindController(db *gorm.DB) *WaterWindController {
	return &WaterWindController{db}
}

func (ww *WaterWindController) Create(c *gin.Context) {
	type RequestBody struct {
		Water int `json:"water" binding:"required"`
		Wind  int `json:"wind" binding:"required"`
	}

	var requestBody RequestBody
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(BadRequestResponse(err))
		return
	}

	waterValue := requestBody.Water
	windValue := requestBody.Wind

	// ●jika water dibawah 5 maka status aman●jika water antara 6 - 8 maka status siaga●jika water diatas 8 maka status bahaya●jika wind dibawah 6 maka status aman●jika wind antara 7 - 15 maka status siaga●jika wind diatas 15 maka status bahaya
	var statusWater string
	if waterValue < 5 {
		statusWater = "Aman"
	} else if waterValue >= 6 && waterValue <= 8 {
		statusWater = "Siaga"
	} else {
		statusWater = "Bahaya"
	}

	var statusWind string
	if windValue < 6 {
		statusWind = "Aman"
	} else if windValue >= 7 && windValue <= 15 {
		statusWind = "Siaga"
	} else {
		statusWind = "Bahaya"
	}

	water := model.Water{
		Value:  waterValue,
		Status: statusWater,
	}

	wind := model.Wind{
		Value:  windValue,
		Status: statusWind,
	}

	err = ww.db.Create(&water).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	err = ww.db.Create(&wind).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(gin.H{
		"water": water,
		"wind":  wind,
	}))

	fmt.Printf("{\n\"water\": %d,\n\"wind\": %d,\n}\nstatus water : %s\nstatus wind : %s\n", waterValue, windValue, statusWater, statusWind)
}
