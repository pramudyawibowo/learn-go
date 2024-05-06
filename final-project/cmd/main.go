package main

import (
	"final-project/internal/controller"
	"final-project/internal/database"
	"final-project/internal/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.Init()
	defer func() {
		dbInstance, _ := db.DB()

		err := dbInstance.Close()
		if err != nil {
			panic("failed to close database")
		}
	}()

	r := gin.Default()

	gin.SetMode(os.Getenv("GIN_MODE"))

	r.Use(gin.Recovery())

	api := r.Group("/api/v1")

	controller.NewAuthController(db).Routes(api)
	controller.NewUserControllers(db).Routes(api, middleware.AuthMiddleware())
	controller.NewRecipientControllers(db).Routes(api, middleware.AuthMiddleware())

	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not Found",
			"data":    nil,
		})
	})

	port := ":" + os.Getenv("APP_PORT")
	r.Run(port)
}
