package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/ume77betty/habit-tracker-be/handlers"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	api := r.Group("/api")
	users := api.Group("/users/:username")

	users.GET("/habits", handlers.GetHabits(db))
	users.POST("/habits", handlers.CreateHabit(db))

	users.GET("/records", handlers.GetRecords(db))
	users.POST("/records", handlers.CreateRecord(db))

	users.GET("/progress", handlers.GetProgress(db))
}
