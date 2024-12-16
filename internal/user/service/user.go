package service

import (
	"errors"

	"CodeLabs/internal/user/repository"
	"CodeLabs/models"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		repo: repository.NewInMemoryUserRepo(),
	}
}

func (s *UserService) Signup(username, email, password string) (bool, error) {
	// Check if user already exists
	if _, err := s.repo.FindByEmail(email); err == nil {
		return false, errors.New("user already exists")
	}

	// Save the new user
	user := models.User{
		Username: username,
		Email:    email,
		Password: password, // Assume plain password for now, we will add validations later
	}
	if err := s.repo.Save(user); err != nil {
		return false, err
	}
	return true, nil
}
