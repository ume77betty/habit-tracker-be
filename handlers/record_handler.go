package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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
