package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ume77betty/habit-tracker-be/models"
	"github.com/ume77betty/habit-tracker-be/services"
)

func GetRecords(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username := ctx.Param("username")
		records, err := services.GetRecordsByUsername(db, username)
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
		ctx.JSON(http.StatusOK, records)
	}

}

func CreateRecord(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.CreateRecordRequest
		username := ctx.Param("username")

		if err := ctx.ShouldBindJSON(&req); err != nil {

			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		result, err := services.CreateRecord(db, username, req)
		if err != nil {

			if errors.Is(err, services.ErrUserNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "user not found",
				})
				return
			}

			if errors.Is(err, services.ErrHabitNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{
					"error": "habit not found",
				})
				return
			}

			if errors.Is(err, services.ErrInvalidRecordTime) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid record time",
				})
				return
			}

			if errors.Is(err, services.ErrInvalidRecordDate) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "invalid record date",
				})
				return
			}

			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, result)
	}
}
