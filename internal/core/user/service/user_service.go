package UserService

import (
	UserModels "go-hexa-full/internal/core/user/models"
	UserPort "go-hexa-full/internal/ports/user"
)

type UserService struct {
	repo UserPort.UserRepository
}

func NewUserService(repo UserPort.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user UserModels.User) error {
	return s.repo.Save(user)
}

func (s *UserService) GetUser(id string) (*UserModels.User, error) {
	return s.repo.FindById(id)
}
