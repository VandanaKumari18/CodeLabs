package repository

import (
	"errors"

	"CodeLabs/models"
)

type UserRepository interface {
	Save(user models.User) error
	FindByEmail(email string) (*models.User, error)
}

type InMemoryUserRepo struct {
	users map[string]models.User
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users: make(map[string]models.User),
	}
}

func (repo *InMemoryUserRepo) Save(user models.User) error {
	if _, exists := repo.users[user.Email]; exists {
		return errors.New("user already exists")
	}
	repo.users[user.Email] = user
	return nil
}

func (repo *InMemoryUserRepo) FindByEmail(email string) (*models.User, error) {
	user, exists := repo.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return &user, nil
}
