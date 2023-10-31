package handler

import (
	"encoding/json"
	"net/http"
	"projeto_chat_backend/pkg/model"
	"projeto_chat_backend/pkg/repository"
	"strconv"

	"github.com/gorilla/mux"
)

// GetTasks recupera todas as tarefas
func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := repository.GetAllTasks()
	if err != nil {
		http.Error(w, "Erro ao buscar tarefas", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// GetTaskByID recupera uma tarefa pelo ID
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	task, err := repository.GetTaskByID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar tarefa", http.StatusInternalServerError)
		return
	}
	if task == nil {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// CreateTask cria uma nova tarefa
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task model.Task

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Erro ao decodificar tarefa", http.StatusBadRequest)
		return
	}

	newTask, err := repository.CreateTask(task)
	if err != nil {
		http.Error(w, "Erro ao criar tarefa", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

// UpdateTask atualiza uma tarefa existente
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var task model.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Erro ao decodificar tarefa", http.StatusBadRequest)
		return
	}
	task.ID = id

	updatedTask, err := repository.UpdateTask(task)
	if err != nil {
		http.Error(w, "Erro ao atualizar tarefa", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTask exclui uma tarefa pelo ID
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := repository.DeleteTask(id); err != nil {
		http.Error(w, "Erro ao excluir tarefa", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
