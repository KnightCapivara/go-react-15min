package main

import (
	"log"
	"net/http"

	"projeto_chat_backend/pkg/handler"
	"projeto_chat_backend/pkg/middleware"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Middlewares
	r.Use(middleware.LoggingMiddleware)       // Suponha que exista um middleware de logging
	r.Use(middleware.ValidateTokenMiddleware) // Middleware de autenticação JWT

	// Endpoints de Usuários
	r.HandleFunc("/users", handler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handler.GetUserByID).Methods("GET")
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", handler.DeleteUser).Methods("DELETE")

	// Endpoints de Projetos
	r.HandleFunc("/projects", handler.GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", handler.GetProjectByID).Methods("GET")
	r.HandleFunc("/projects", handler.CreateProject).Methods("POST")
	r.HandleFunc("/projects/{id}", handler.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", handler.DeleteProject).Methods("DELETE")

	// ... Adicione outros endpoints conforme necessário ...

	// Inicie o servidor na porta 8080
	log.Println("Servidor rodando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
