package repository

import (
	"projeto_chat_backend/pkg/config"
	"projeto_chat_backend/pkg/model"
)

// GetAllTasks recupera todas as tarefas do banco de dados
func GetAllTasks() ([]model.Task, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, description, status, project_id FROM Tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Status, &t.ProjectID); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// GetTaskByID recupera uma tarefa específica pelo ID
func GetTaskByID(id int) (*model.Task, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var t model.Task
	err = db.QueryRow("SELECT id, name, description, status, project_id FROM Tasks WHERE id = $1", id).Scan(&t.ID, &t.Name, &t.Description, &t.Status, &t.ProjectID)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

// CreateTask insere uma nova tarefa no banco de dados
func CreateTask(task model.Task) (*model.Task, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.QueryRow("INSERT INTO Tasks (name, description, status, project_id) VALUES ($1, $2, $3, $4) RETURNING id", task.Name, task.Description, task.Status, task.ProjectID).Scan(&task.ID)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateTask atualiza uma tarefa existente no banco de dados
func UpdateTask(task model.Task) (*model.Task, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Tasks SET name = $1, description = $2, status = $3, project_id = $4 WHERE id = $5", task.Name, task.Description, task.Status, task.ProjectID, task.ID)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// DeleteTask exclui uma tarefa pelo ID
func DeleteTask(id int) error {
	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Tasks WHERE id = $1", id)
	return err
}
