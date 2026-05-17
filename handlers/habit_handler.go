package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ume77betty/habit-tracker-be/models"
	"github.com/ume77betty/habit-tracker-be/services"
)

func GetHabits(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		habits, err := services.GetHabitsByUsername(db, username)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "user not found",
				})
				return
			}
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}
		ctx.JSON(http.StatusOK, habits)
	}

}

func CreateHabit(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.CreateHabitRequest

		if err := ctx.ShouldBindJSON(&req); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		username := ctx.Param("username")
		result, err := services.CreateHabit(db, username, req)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "user not found",
				})
				return
			}

			if errors.Is(err, services.ErrInvalidName) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid name",
				})
				return
			}

			if errors.Is(err, services.ErrInvalidTargetDays) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid target days",
				})
				return
			}

			if errors.Is(err, services.ErrDuplicateName) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "duplicate habit name",
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		ctx.JSON(http.StatusCreated, result)
	}
}
