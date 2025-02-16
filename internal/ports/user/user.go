package userport

import UserModels "go-hexa-full/internal/core/user/models"

type UserRepository interface {
	Save(user UserModels.User) error
	FindById(id string) (*UserModels.User, error)
}
