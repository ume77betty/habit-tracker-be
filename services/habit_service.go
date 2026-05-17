package services

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/ume77betty/habit-tracker-be/models"
	"github.com/ume77betty/habit-tracker-be/repositories"
)

var ErrInvalidName = errors.New("name cannot be empty")
var ErrInvalidTargetDays = errors.New("target days should larger than zero")
var ErrDuplicateName = errors.New("duplicate habit name")

func GetHabitsByUsername(db *sql.DB, username string) ([]models.Habit, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return nil, err
	}

	return repositories.GetHabitsByUserID(db, user.ID)
}

func CreateHabit(db *sql.DB, username string, req models.CreateHabitRequest) (models.CreateHabitResponse, error) {
	if req.TargetDays <= 0 {
		return models.CreateHabitResponse{}, ErrInvalidTargetDays
	}

	if len(strings.TrimSpace(req.Name)) <= 0 {
		return models.CreateHabitResponse{}, ErrInvalidName
	}

	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return models.CreateHabitResponse{}, err
	}

	exists, err := repositories.CheckDuplicateHabitName(db, user.ID, req.Name)

	if err != nil {
		return models.CreateHabitResponse{}, err
	}

	if exists == true {
		return models.CreateHabitResponse{}, ErrDuplicateName
	}

	habit, err := repositories.CreateHabit(db, user.ID, req)
	if err != nil {
		return models.CreateHabitResponse{}, err
	}

	return models.CreateHabitResponse{ID: habit.ID}, nil
}
