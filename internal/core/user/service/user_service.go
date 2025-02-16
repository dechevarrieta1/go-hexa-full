package userservice

import (
	"errors"
	userModels "go-hexa-full/internal/core/user/models"
	userPort "go-hexa-full/internal/ports/user"
)

type UserService struct {
	repo userPort.UserRepository
}

func NewUserService(repo userPort.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user userModels.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("missing required fields")
	}

	existingUser, _ := s.repo.FindById(user.Name)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	return s.repo.Save(user)
}

func (s *UserService) GetUser(id string) (*userModels.User, error) {
	if id == "" {
		return nil, errors.New("user ID is required")
	}

	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
