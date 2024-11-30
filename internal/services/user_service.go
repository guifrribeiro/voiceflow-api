package services

import (
	"errors"
	"voiceflow/internal/models"
	"voiceflow/internal/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("nome e e-mail são obrigatórios")
	}

	return s.UserRepo.CreateUser(user)
}