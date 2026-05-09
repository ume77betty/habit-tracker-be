package repositories

import "database/sql"

func GetWeeklyProgressByUserID(db *sql.DB, userID string, startDate string, endDate string) (map[string]int, error) {
	query := `
		SELECT habit_id, COUNT(*)
		FROM records
		WHERE user_id = $1
			AND record_date >= $2
			AND record_date <= $3
			AND is_deleted = FALSE
		GROUP BY habit_id
	`

	rows, err := db.Query(query, userID, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	progress := map[string]int{}

	for rows.Next() {
		var habitID string
		var count int

		err := rows.Scan(&habitID, &count)
		if err != nil {
			return nil, err
		}

		progress[habitID] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return progress, nil
}
