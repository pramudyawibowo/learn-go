package controller

import (
	"final-project/internal/helper"
	"final-project/internal/model"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{db}
}

func (a *AuthController) Routes(r *gin.RouterGroup) {
	r.POST("/login", a.Login)
	r.POST("/register", a.Register)
}

func (a *AuthController) Login(c *gin.Context) {
	var LoginRequest model.LoginRequest

	err := c.ShouldBindJSON(&LoginRequest)
	if err != nil {
		c.JSON(BadRequestResponse(err))
		return
	}

	// using govalidator
	_, err = govalidator.ValidateStruct(&LoginRequest)
	if err != nil {
		c.JSON(BadRequestResponse(err))
		return
	}

	var user model.User
	err = a.db.Where("username = ?", LoginRequest.Username).First(&user).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	err = helper.ComparePassword(user.Password, LoginRequest.Password)
	if err != nil {
		c.JSON(UnauthorizedResponse(err))
		return
	}

	user.Token, err = helper.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	err = a.db.Save(&user).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(gin.H{
		"token": user.Token,
	}))
}

func (a *AuthController) Register(c *gin.Context) {
	var CreateUserRequest model.User

	err := c.ShouldBindJSON(&CreateUserRequest)
	if err != nil {
		c.JSON(BadRequestResponse(err))
		return
	}

	err = a.db.Create(&CreateUserRequest).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(CreateUserRequest))
}
