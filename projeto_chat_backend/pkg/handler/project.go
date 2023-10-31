package handler

import (
	"encoding/json"
	"net/http"
	"projeto_chat_backend/pkg/model"
	"projeto_chat_backend/pkg/repository"
	"strconv"

	"github.com/gorilla/mux"
)

// GetProjects recupera todos os projetos
func GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := repository.GetAllProjects()
	if err != nil {
		http.Error(w, "Erro ao buscar projetos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// GetProjectByID recupera um projeto pelo ID
func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	project, err := repository.GetProjectByID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar projeto", http.StatusInternalServerError)
		return
	}
	if project == nil {
		http.Error(w, "Projeto não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

// CreateProject cria um novo projeto
func CreateProject(w http.ResponseWriter, r *http.Request) {
	var project model.Project

	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Erro ao decodificar projeto", http.StatusBadRequest)
		return
	}

	newProject, err := repository.CreateProject(project)
	if err != nil {
		http.Error(w, "Erro ao criar projeto", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProject)
}

// UpdateProject atualiza um projeto existente
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var project model.Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Erro ao decodificar projeto", http.StatusBadRequest)
		return
	}
	project.ID = id

	updatedProject, err := repository.UpdateProject(project)
	if err != nil {
		http.Error(w, "Erro ao atualizar projeto", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProject)
}

// DeleteProject exclui um projeto pelo ID
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := repository.DeleteProject(id); err != nil {
		http.Error(w, "Erro ao excluir projeto", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
