package compositeadapter

import (
	usermodels "go-hexa-full/internal/core/user/models"
	userport "go-hexa-full/internal/ports/user"
)

type CompositeUserRepository struct {
	repositories []userport.UserRepository
}

func NewCompositeUserRepository(repos ...userport.UserRepository) *CompositeUserRepository {
	return &CompositeUserRepository{repositories: repos}
}

func (r *CompositeUserRepository) Save(user usermodels.User) error {
	for _, repo := range r.repositories {
		if err := repo.Save(user); err != nil {
			return err
		}
	}
	return nil
}

func (r *CompositeUserRepository) FindById(id string) (*usermodels.User, error) {
	for _, repo := range r.repositories {
		user, err := repo.FindById(id)
		if err != nil {
			return nil, err
		}
		if user != nil {
			return user, nil
		}
	}
	return nil, nil
}
