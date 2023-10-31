package repository

import (
	"encoding/json"
	"net/http"
	"projeto_chat_backend/pkg/auth"
	"projeto_chat_backend/pkg/config"
	"projeto_chat_backend/pkg/model"
	"time"
)

func StoreRefreshToken(token model.RefreshToken) error {
	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO refresh_tokens (token, user_id, expiry, created_at) VALUES ($1, $2, $3, $4)", token.Token, token.UserID, token.Expiry, token.CreatedAt)
	return err
}

func RetrieveRefreshToken(tokenString string) (*model.RefreshToken, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var token model.RefreshToken
	err = db.QueryRow("SELECT token, user_id, expiry, created_at FROM refresh_tokens WHERE token = $1", tokenString).Scan(&token.Token, &token.UserID, &token.Expiry, &token.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func DeleteRefreshToken(tokenString string) error {
	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM refresh_tokens WHERE token = $1", tokenString)
	return err
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

	json.NewEncoder(w).Encode(map[string]string{"token": newAccessToken})
}

func Revoke(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Refresh-Token")

	// Excluir o refresh token do banco de dados
	err := repository.DeleteRefreshToken(tokenString)
	if err != nil {
		http.Error(w, "Erro ao revogar refresh token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
