package controller

import (
	"final-project/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserControllers struct {
	db *gorm.DB
}

func NewUserControllers(db *gorm.DB) *UserControllers {
	return &UserControllers{db}
}

func (u *UserControllers) Routes(r *gin.RouterGroup, middlewares ...gin.HandlerFunc) {

	r.Use(middlewares...)

	r.GET("/users", u.GetAllUser)
	r.GET("/users/:id", u.GetUserByID)
	r.POST("/users", u.Create)
	// r.PUT("/users/:id", u.Update)
	r.DELETE("/users/:id", u.Delete)
}

func (u *UserControllers) GetAllUser(c *gin.Context) {
	var users []model.User

	err := u.db.Find(&users).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(users))
}

func (u *UserControllers) GetUserByID(c *gin.Context) {
	var user model.User

	err := u.db.First(&user, c.Param("id")).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(user))
}

func (u *UserControllers) Create(c *gin.Context) {
	var CreateUserRequest model.User

	err := c.ShouldBindJSON(&CreateUserRequest)
	if err != nil {
		c.JSON(BadRequestResponse(err))
		return
	}

	err = u.db.Create(&CreateUserRequest).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(CreateUserRequest))
}

// func (u *UserControllers) Update(c *gin.Context) {
// 	var user model.User

// 	err := u.db.First(&user, c.Param("id")).Error
// 	if err != nil {
// 		c.JSON(InternalServerErrorResponse(err))
// 		return
// 	}

// 	var UpdateUserRequest model.UpdateUserRequest

// 	err = c.ShouldBindJSON(&UpdateUserRequest)
// 	if err != nil {
// 		c.JSON(BadRequestResponse(err))
// 		return
// 	}

// 	if UpdateUserRequest.Name != nil {
// 		user.Name = *UpdateUserRequest.Name
// 	}

// 	if UpdateUserRequest.Username != nil {
// 		user.Username = *UpdateUserRequest.Username
// 	}

// 	if UpdateUserRequest.Password != nil {
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*UpdateUserRequest.Password), bcrypt.DefaultCost)
// 		if err != nil {
// 			c.JSON(InternalServerErrorResponse(err))
// 			return
// 		}

// 		user.Password = hashedPassword
// 	}

// 	if UpdateUserRequest.Email != nil {
// 		user.Email = *UpdateUserRequest.Email
// 	}

// 	if UpdateUserRequest.Role != nil {
// 		user.Role = *UpdateUserRequest.Role
// 	}

// 	err = u.db.Save(&user).Error
// 	if err != nil {
// 		c.JSON(InternalServerErrorResponse(err))
// 		return
// 	}

// 	c.JSON(SuccessResponse(user))
// }

func (u *UserControllers) Delete(c *gin.Context) {
	var user model.User

	err := u.db.First(&user, c.Param("id")).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	err = u.db.Delete(&user).Error
	if err != nil {
		c.JSON(InternalServerErrorResponse(err))
		return
	}

	c.JSON(SuccessResponse(nil))
}
