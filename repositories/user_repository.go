package repositories

import (
	"database/sql"

	"github.com/ume77betty/habit-tracker-be/models"
)

func GetUserByUsername(db *sql.DB, username string) (*models.User, error) {
	var user models.User

	query := `
	SELECT id, username, created_at, updated_at FROM users WHERE username = $1
	`

	err := db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
