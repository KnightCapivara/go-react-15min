package model

// Project representa a entidade de um projeto
type Project struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
}
