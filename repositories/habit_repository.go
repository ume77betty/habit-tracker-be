package repositories

import (
	"database/sql"

	"github.com/ume77betty/habit-tracker-be/models"
)

func GetHabitsByUserID(db *sql.DB, userID string) ([]models.Habit, error) {
	query := `
		SELECT id, name, icon_name, color, target_days, last_recorded_at, created_at, updated_at, is_active
		FROM habits
		WHERE user_id = $1
		ORDER BY created_at ASC
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	habits := []models.Habit{}

	for rows.Next() {
		var habit models.Habit
		var lastRecordedAt sql.NullTime
		var createdAt sql.NullTime
		var updatedAt sql.NullTime

		err := rows.Scan(
			&habit.ID,
			&habit.Name,
			&habit.IconName,
			&habit.Color,
			&habit.TargetDays,
			&lastRecordedAt,
			&createdAt,
			&updatedAt,
			&habit.IsActive,
		)
		if err != nil {
			return nil, err
		}

		if lastRecordedAt.Valid {
			habit.LastRecordedAt = &lastRecordedAt.Time
		}
		if createdAt.Valid {
			habit.CreatedAt = &createdAt.Time
		}
		if updatedAt.Valid {
			habit.UpdatedAt = &updatedAt.Time
		}

		habits = append(habits, habit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return habits, nil
}
