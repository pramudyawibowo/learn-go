package controller

import (
	"final-project/internal/helper"
	"final-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RecipientControllers struct {
	db *gorm.DB
}

func NewRecipientControllers(db *gorm.DB) *RecipientControllers {
	return &RecipientControllers{db}
}

func (rc *RecipientControllers) Routes(r *gin.RouterGroup, middlewares ...gin.HandlerFunc) {
	r.GET("recipients", rc.Get)
	r.GET("recipients/:id", rc.Show)
	r.POST("recipients", rc.Create)
	r.PUT("recipients/:id", rc.Update)
	r.DELETE("recipients/:id", rc.Delete)
}

func (r *RecipientControllers) Get(c *gin.Context) {
	user, _ := c.MustGet("user").(*helper.JWTClaims)

	var recipients []model.Recipient
	err := r.db.Where("user_id = ?", user.ID).Find(&recipients).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(recipients))
}

func (r *RecipientControllers) Show(c *gin.Context) {
	user, _ := c.MustGet("user").(*helper.JWTClaims)

	var recipientData model.Recipient
	err := r.db.Where("id = ? AND user_id = ?", c.Param("id"), user.ID).First(&recipientData).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(recipientData))
}

func (r *RecipientControllers) Create(c *gin.Context) {
	var recipient model.CreateRecipientRequest

	err := c.ShouldBindJSON(&recipient)
	if err != nil {
		c.JSON(BadRequestResponse(err))
		return
	}

	user, _ := c.MustGet("user").(*helper.JWTClaims)

	newRecipient := model.Recipient{
		UserID:      user.ID,
		Name:        recipient.Name,
		Phonenumber: recipient.Phonenumber,
		Address:     recipient.Address,
		City:        recipient.City,
		Province:    recipient.Province,
	}

	err = r.db.Create(&newRecipient).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(newRecipient))
}

func (r *RecipientControllers) Update(c *gin.Context) {
	var recipient model.UpdateRecipientRequest

	err := c.ShouldBindJSON(&recipient)
	if err != nil {
		c.JSON(BadRequestResponse(err))
		return
	}

	user, _ := c.MustGet("user").(*helper.JWTClaims)

	var recipientData model.Recipient
	err = r.db.Where("id = ? AND user_id = ?", c.Param("id"), user.ID).First(&recipientData).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	if recipient.Name != nil {
		recipientData.Name = *recipient.Name
	}

	if recipient.Phonenumber != nil {
		recipientData.Phonenumber = *recipient.Phonenumber
	}

	if recipient.Address != nil {
		recipientData.Address = *recipient.Address
	}

	if recipient.City != nil {
		recipientData.City = *recipient.City
	}

	if recipient.Province != nil {
		recipientData.Province = *recipient.Province
	}

	err = r.db.Save(&recipientData).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(recipientData))
}

func (r *RecipientControllers) Delete(c *gin.Context) {
	user, _ := c.MustGet("user").(*helper.JWTClaims)

	var recipientData model.Recipient
	err := r.db.Where("id = ? AND user_id = ?", c.Param("id"), user.ID).First(&recipientData).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	err = r.db.Delete(&recipientData).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(nil))
}
