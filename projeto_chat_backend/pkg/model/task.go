package model

// Task representa a entidade de uma tarefa
type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ProjectID   int    `json:"project_id"`
}
