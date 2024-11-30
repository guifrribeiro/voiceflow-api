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
	query := "INSERT INTO voiceflow.users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
	err := repo.DB.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.Id)

	if err != nil {
		return errors.New("Erro ao cadastrar usuário: " + err.Error())
	}

	return nil
}

func (repo *UserRepository) GetUsers() ([]models.User, error) {
	query := "SELECT * FROM voiceflow.users"
	rows, err := repo.DB.Query(query)

	if err != nil {
		return nil, errors.New("erro ao buscar usuários: " + err.Error())
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}