package repository

import (
	"projeto_chat_backend/pkg/config"
	"projeto_chat_backend/pkg/model"
)

// GetAllProjects recupera todos os projetos do banco de dados
func GetAllProjects() ([]model.Project, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, description, user_id FROM Projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []model.Project
	for rows.Next() {
		var p model.Project
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.UserID); err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}
	return projects, nil
}

// GetProjectByID recupera um projeto específico pelo ID
func GetProjectByID(id int) (*model.Project, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var p model.Project
	err = db.QueryRow("SELECT id, name, description, user_id FROM Projects WHERE id = $1", id).Scan(&p.ID, &p.Name, &p.Description, &p.UserID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// CreateProject insere um novo projeto no banco de dados
func CreateProject(project model.Project) (*model.Project, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.QueryRow("INSERT INTO Projects (name, description, user_id) VALUES ($1, $2, $3) RETURNING id", project.Name, project.Description, project.UserID).Scan(&project.ID)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// UpdateProject atualiza um projeto existente no banco de dados
func UpdateProject(project model.Project) (*model.Project, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Projects SET name = $1, description = $2, user_id = $3 WHERE id = $4", project.Name, project.Description, project.UserID, project.ID)
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// DeleteProject exclui um projeto pelo ID
func DeleteProject(id int) error {
	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Projects WHERE id = $1", id)
	return err
}
