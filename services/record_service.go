package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/ume77betty/habit-tracker-be/models"
	"github.com/ume77betty/habit-tracker-be/repositories"
)

var ErrUserNotFound = errors.New("user not found")
var ErrHabitNotFound = errors.New("habit not found")
var ErrInvalidRecordTime = errors.New("invalid record time")
var ErrInvalidRecordDate = errors.New("invalid record date")

func GetRecordsByUsername(db *sql.DB, username string) ([]models.Record, error) {
	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		return nil, err
	}
	_ = user
	return repositories.GetRecordsByUserID(db, user.ID)
}

func CreateRecord(db *sql.DB, username string, req models.CreateRecordRequest) (models.CreateRecordResponse, error) {
	startTime, err := time.Parse(time.RFC3339, req.StartTime)
	if err != nil {
		return models.CreateRecordResponse{}, ErrInvalidRecordTime
	}

	endTime, err := time.Parse(time.RFC3339, req.EndTime)
	if err != nil {
		return models.CreateRecordResponse{}, ErrInvalidRecordTime
	}
	if endTime.Before(startTime) {
		return models.CreateRecordResponse{}, ErrInvalidRecordTime
	}

	_, err = time.Parse("2006-01-02", req.RecordDate)
	if err != nil {
		return models.CreateRecordResponse{}, ErrInvalidRecordDate
	}

	user, err := repositories.GetUserByUsername(db, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.CreateRecordResponse{}, ErrUserNotFound
		}
		return models.CreateRecordResponse{}, err
	}

	_, err = repositories.GetHabitByIDAndUserID(db, req.HabitID, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.CreateRecordResponse{}, ErrHabitNotFound
		}
		return models.CreateRecordResponse{}, err
	}

	record, err := repositories.CreateRecord(db, user.ID, req)
	if err != nil {
		return models.CreateRecordResponse{}, err
	}

	err = repositories.UpdateHabitLastRecordedAt(db, req.HabitID, startTime)
	if err != nil {
		return models.CreateRecordResponse{}, err
	}

	return models.CreateRecordResponse{
		Message:  "record payload received",
		Username: username,
		UserID:   user.ID,
		Record:   record,
	}, nil
}
