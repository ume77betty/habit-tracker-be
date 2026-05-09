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
