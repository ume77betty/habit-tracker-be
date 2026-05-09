package main

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ume77betty/habit-tracker-be/db"
	"github.com/ume77betty/habit-tracker-be/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	database := db.NewDB()
	defer database.Close()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(r, database)

	r.Run() // 預設監聽 0.0.0.0:8080
}
