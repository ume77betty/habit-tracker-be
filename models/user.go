package models

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
