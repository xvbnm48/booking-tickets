package service

import (
	"booking-ticket/internal/model"
	"booking-ticket/internal/repository"
	"booking-ticket/logger"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAllUsers() ([]model.UserResponse, error)
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.CreateUserRequest) (bool, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsers() ([]model.UserResponse, error) {
	return s.repo.FindAllUsers()
}

func (s *userService) GetUserByEmail(email string) (model.User, error) {
	result, err := s.repo.FindUserByEmail(email)
	fmt.Println("isi error", err)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return model.User{}, fmt.Errorf("user not found")
		}
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return result, nil
}

func (s *userService) CreateUser(user model.CreateUserRequest) (bool, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}
	user.Password = string(hashedPassword)

	exiestEmail, err := s.repo.CheckEmailExists(user.Email)
	if err != nil {
		logger.Error("failed to check email exists", err)
		return false, err
	}

	if exiestEmail {
		return false, fmt.Errorf("email already exists")
	}
	newUser, err := s.repo.CreateUser(user)
	if err != nil {
		return false, err
	}
	logger.Info("new user created")
	return newUser, nil
}
