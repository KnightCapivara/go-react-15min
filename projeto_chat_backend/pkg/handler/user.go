package handler

import (
	"encoding/json"
	"net/http"
	"projeto_chat_backend/internal/util"
	"projeto_chat_backend/pkg/auth"
	"projeto_chat_backend/pkg/model"
	"projeto_chat_backend/pkg/repository"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// GetUsers recupera todos os usuários
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repository.GetAllUsers()
	if err != nil {
		http.Error(w, "Erro ao buscar usuários", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// GetUserByID recupera um usuário pelo ID
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	user, err := repository.GetUserByID(id)
	if err != nil {
		http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
		return
	}
	if user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUser cria um novo usuário
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar usuário", http.StatusBadRequest)
		return
	}

	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "Erro ao hashear senha", http.StatusInternalServerError)
		return
	}
	user.Password = hashedPassword

	newUser, err := repository.CreateUser(user)
	if err != nil {
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

// UpdateUser atualiza um usuário existente
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Erro ao decodificar usuário", http.StatusBadRequest)
		return
	}
	user.ID = id

	if user.Password != "" {
		hashedPassword, err := util.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Erro ao hashear senha", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword
	}

	updatedUser, err := repository.UpdateUser(user)
	if err != nil {
		http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser exclui um usuário pelo ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if err := repository.DeleteUser(id); err != nil {
		http.Error(w, "Erro ao excluir usuário", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func Login(w http.ResponseWriter, r *http.Request) {
	// ... [código existente]

	refreshToken, err := auth.GenerateRefreshToken(*user)
	if err != nil {
		http.Error(w, "Erro ao gerar refresh token", http.StatusInternalServerError)
		return
	}

	// Armazene o refresh token no banco de dados
	err = repository.StoreRefreshToken(*refreshToken)
	if err != nil {
		http.Error(w, "Erro ao armazenar refresh token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token":        tokenString,
		"refreshToken": refreshToken.Token,
	})
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Refresh-Token")

	// Recuperar o refresh token do banco de dados
	storedToken, err := repository.RetrieveRefreshToken(tokenString)
	if err != nil {
		http.Error(w, "Refresh token inválido", http.StatusUnauthorized)
		return
	}

	// Verificar a validade do refresh token
	if time.Now().After(storedToken.Expiry) {
		http.Error(w, "Refresh token expirado", http.StatusUnauthorized)
		return
	}

	user, err := repository.GetUserByID(storedToken.UserID)
	if err != nil || user == nil {
		http.Error(w, "Erro ao recuperar o usuário", http.StatusInternalServerError)
		return
	}

	// Gerar um novo access token
	newAccessToken, err := auth.GenerateToken(*user)
	if err != nil {
		http.Error(w, "Erro ao gerar access token", http.StatusInternalServerError)
		return
	}

	// Gerar um novo refresh token
	newRefreshToken, err := auth.GenerateRefreshToken(*user) // Esta é uma nova função que você precisa criar
	if err != nil {
		http.Error(w, "Erro ao gerar refresh token", http.StatusInternalServerError)
		return
	}

	// Armazenar o novo refresh token no banco de dados e remover o antigo
	err = repository.StoreRefreshToken(newRefreshToken)
	if err != nil {
		http.Error(w, "Erro ao armazenar refresh token", http.StatusInternalServerError)
		return
	}
	err = repository.DeleteRefreshToken(tokenString)
	if err != nil {
		http.Error(w, "Erro ao remover o antigo refresh token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token":        newAccessToken,
		"refreshToken": newRefreshToken.Token,
	})
}
