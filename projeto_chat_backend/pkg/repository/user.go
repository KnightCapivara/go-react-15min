package repository

import (
	"database/sql"
	"projeto_chat_backend/pkg/config"
	"projeto_chat_backend/pkg/model"
)

// GetAllUsers recupera todos os usuários do banco de dados
func GetAllUsers() ([]model.User, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, username FROM Users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		if err := rows.Scan(&u.ID, &u.Username); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

// GetUserByID recupera um usuário pelo ID
func GetUserByID(id string) (*model.User, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var u model.User
	err = db.QueryRow("SELECT id, username FROM Users WHERE id = $1", id).Scan(&u.ID, &u.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}

// CreateUser insere um novo usuário no banco de dados
func CreateUser(user model.User) (*model.User, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	err = db.QueryRow("INSERT INTO Users (username, password) VALUES ($1, $2) RETURNING id", user.Username, user.Password).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser atualiza um usuário no banco de dados
func UpdateUser(user model.User) (*model.User, error) {
	db, err := config.GetDBConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("UPDATE Users SET username = $1, password = $2 WHERE id = $3", user.Username, user.Password, user.ID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// DeleteUser exclui um usuário pelo ID
func DeleteUser(id string) error {
	db, err := config.GetDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Users WHERE id = $1", id)
	return err
}
