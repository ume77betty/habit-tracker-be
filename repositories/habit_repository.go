package repositories

import (
	"database/sql"
	"time"

	"github.com/ume77betty/habit-tracker-be/models"
)

func UpdateHabitLastRecordedAt(tx *sql.Tx, habitID string, recordedAt time.Time) error {
	query := `
		UPDATE habits
		SET last_recorded_at = $1, updated_at = now()
		WHERE id = $2
			AND (last_recorded_at IS NULL OR last_recorded_at < $1)
	`
	_, err := tx.Exec(query, recordedAt, habitID)
	if err != nil {
		return err
	}
	return nil
}

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

func GetHabitByIDAndUserID(db *sql.DB, habitID string, userID string) (*models.Habit, error) {
	var habit models.Habit
	query := `
		SELECT id, name, icon_name, color, target_days, last_recorded_at, created_at, updated_at, is_active
		FROM habits
		WHERE id = $1 AND user_id = $2
	`

	var lastRecordedAt sql.NullTime
	var createdAt sql.NullTime
	var updatedAt sql.NullTime

	err := db.QueryRow(query, habitID, userID).Scan(
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

	return &habit, nil
}

func CreateHabit(db *sql.DB, userID string, req models.CreateHabitRequest) (models.CreatedHabit, error) {
	var habit models.CreatedHabit
	query := `
		INSERT INTO habits (
			user_id,
			name,
			icon_name,
			color,
			target_days
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err := db.QueryRow(
		query,
		userID,
		req.Name,
		req.IconName,
		req.Color,
		req.TargetDays,
	).Scan(&habit.ID)

	if err != nil {
		return models.CreatedHabit{}, err
	}
	return habit, nil
}
