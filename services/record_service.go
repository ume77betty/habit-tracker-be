package services

import (
	"database/sql"

	"github.com/ume77betty/habit-tracker-be/models"
	"github.com/ume77betty/habit-tracker-be/repositories"
)

func GetRecordsByUsername(db *sql.DB, username string) ([]models.Record, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return nil, err
	}
	_ = user
	return repositories.GetRecordsByUserID(db, user.ID)
}
