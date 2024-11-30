package repositories

import (
	"database/sql"
	"errors"
	"voiceflow/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	err := repo.DB.QueryRow(query, user.Name, user.Email).Scan(&user.Id)

	if err != nil {
		return errors.New("Erro ao cadastrar usuário: " + err.Error())
	}

	return nil
}