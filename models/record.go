package models

import "time"

type Record struct {
	ID         string     `json:"id"`
	HabitID    string     `json:"habitId"`
	RecordDate string     `json:"recordDate"`
	StartTime  time.Time  `json:"startTime"`
	EndTime    time.Time  `json:"endTime"`
	Reflection string     `json:"reflection"`
	MoodLevel  int        `json:"moodLevel"`
	TZ         string     `json:"tz"`
	IsDeleted  bool       `json:"isDeleted"`
	CreatedAt  *time.Time `json:"createdAt,omitempty"`
	UpdatedAt  *time.Time `json:"updatedAt,omitempty"`
}

type CreateRecordRequest struct {
	StartTime  string `json:"startTime" binding:"required"`
	EndTime    string `json:"endTime" binding:"required"`
	Reflection string `json:"reflection"`
	MoodLevel  int    `json:"moodLevel" binding:"required"`
	HabitID    string `json:"habitId" binding:"required"`
	RecordDate string `json:"recordDate" binding:"required"`
	TZ         string `json:"tz" binding:"required"`
}

type CreatedRecord struct {
	ID string `json:"id"`
}

type CreateRecordResponse struct {
	Message  string        `json:"message"`
	Username string        `json:"username"`
	UserID   string        `json:"userId"`
	Record   CreatedRecord `json:"record"`
}
