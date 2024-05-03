package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sesi7/internal/controller"
	"sesi7/internal/database"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type WaterWindData struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

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

	r.Use(gin.Recovery())

	waterWindController := controller.NewWaterWindController(db)
	r.POST("/waterwind", waterWindController.Create)

	r.Any("/", func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Not Found",
			"data":    nil,
		})
	})

	port := ":" + os.Getenv("APP_PORT")
	go func() {
		for {
			time.Sleep(3 * time.Second)

			water := rand.Intn(100) + 1
			wind := rand.Intn(100) + 1

			data := WaterWindData{
				Water: water,
				Wind:  wind,
			}
			payload, err := json.Marshal(data)
			if err != nil {
				log.Println("Error marshalling JSON:", err)
				continue
			}

			_, err = http.Post("http://localhost"+port+"/waterwind", "application/json", bytes.NewBuffer(payload))
			if err != nil {
				log.Println("Failed to hit /waterwind endpoint:", err)
			}
		}
	}()

	r.Run(port)
}
