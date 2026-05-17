package models

import "time"

type Habit struct {
	ID             string     `json:"id"`
	Name           string     `json:"name"`
	IconName       string     `json:"iconName"`
	Color          string     `json:"color"`
	TargetDays     int        `json:"targetDays"`
	LastRecordedAt *time.Time `json:"lastRecordedAt"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
	UpdatedAt      *time.Time `json:"updatedAt,omitempty"`
	IsActive       bool       `json:"isActive"`
}

type CreateHabitRequest struct {
	Name       string `json:"name" binding:"required"`
	IconName   string `json:"iconName" binding:"required"`
	Color      string `json:"color" binding:"required"`
	TargetDays int    `json:"targetDays" binding:"required"`
}

type CreateHabitResponse struct {
	ID string `json:"id"`
}

type CreatedHabit struct {
	ID string `json:"id"`
}
