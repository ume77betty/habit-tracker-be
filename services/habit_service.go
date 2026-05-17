package services

import (
	"database/sql"

	"github.com/ume77betty/habit-tracker-be/models"
	"github.com/ume77betty/habit-tracker-be/repositories"
)

func GetHabitsByUsername(db *sql.DB, username string) ([]models.Habit, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return nil, err
	}

	return repositories.GetHabitsByUserID(db, user.ID)
}

func CreateHabit(db *sql.DB, username string, req models.CreateHabitRequest) (models.CreateHabitResponse, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return models.CreateHabitResponse{}, err
	}

	habit, err := repositories.CreateHabit(db, user.ID, req)
	if err != nil {
		return models.CreateHabitResponse{}, err
	}

	return models.CreateHabitResponse{ID: habit.ID}, nil
}
