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
