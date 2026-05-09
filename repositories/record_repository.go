package repositories

import (
	"database/sql"
	"time"

	"github.com/ume77betty/habit-tracker-be/models"
)

func GetRecordsByUserID(db *sql.DB, userID string) ([]models.Record, error) {
	query := `
		SELECT id, habit_id, record_date, start_time, end_time, reflection, mood_level, tz, is_deleted, created_at, updated_at
		FROM records
		WHERE user_id = $1
		ORDER BY record_date DESC, created_at DESC
	`

	rows, err := db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	records := []models.Record{}

	for rows.Next() {
		var record models.Record
		var recordDate time.Time
		var createdAt sql.NullTime
		var updatedAt sql.NullTime

		err := rows.Scan(
			&record.ID,
			&record.HabitID,
			&recordDate,
			&record.StartTime,
			&record.EndTime,
			&record.Reflection,
			&record.MoodLevel,
			&record.TZ,
			&record.IsDeleted,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			return nil, err
		}

		record.RecordDate = recordDate.Format("2006-01-02")
		if createdAt.Valid {
			record.CreatedAt = &createdAt.Time
		}
		if updatedAt.Valid {
			record.UpdatedAt = &updatedAt.Time
		}
		records = append(records, record)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
