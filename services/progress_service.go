package services

import (
	"database/sql"
	"time"

	"github.com/ume77betty/habit-tracker-be/models"
	"github.com/ume77betty/habit-tracker-be/repositories"
)

func GetWeeklyProgressByUsername(db *sql.DB, username string) (models.WeeklyProgress, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	weekday := int(now.Weekday())
	startOfWeek := now.AddDate(0, 0, -weekday)

	startDate := startOfWeek.Format("2006-01-02")
	endDate := now.Format("2006-01-02")

	progress, err := repositories.GetWeeklyProgressByUserID(db, user.ID, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return models.WeeklyProgress(progress), nil
}
