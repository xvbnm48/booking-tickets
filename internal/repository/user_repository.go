package repository

import (
	"booking-ticket/internal/model"
	"booking-ticket/logger"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

type UserRepository interface {
	FindAllUsers() ([]model.UserResponse, error)
	FindUserByEmail(email string) (model.User, error)
	CreateUser(user model.CreateUserRequest) (bool, error)
	CheckEmailExists(email string) (bool, error)
}

type userRepository struct {
	db *sql.DB
}

var ErrUserNotFound = errors.New("user not found")

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindAllUsers() ([]model.UserResponse, error) {
	logger.Info("Get All Users")
	rows, err := r.db.Query("SELECT id, name, email, phone_number FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []model.UserResponse
	for rows.Next() {
		var user model.UserResponse
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.PhoneNumber)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) FindUserByEmail(email string) (model.User, error) {
	logger.Info("Find User By Email")
	query := "SELECT id, name, email, password, phone_number FROM users WHERE email = ?"
	log.Println("isi query", query)
	row := r.db.QueryRow(query, email)
	var user model.User

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.PhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}

	return user, nil
}

func (r *userRepository) CreateUser(user model.CreateUserRequest) (bool, error) {
	logger.Info("Create new user")
	query := "INSERT INTO users (name, email, password, phone_number) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, user.Name, user.Email, user.Password, user.PhoneNumber)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rowsAffected == 0 {
		return false, errors.New("no rows affected")
	}
	fmt.Println("isi rows affected", rowsAffected)

	return true, nil

}

func (s *userRepository) CheckEmailExists(email string) (bool, error) {
	logger.Info("Check Email Exists")
	var count int
	query := "SELECT COUNT(*) FROM users WHERE email = ?"
	err := s.db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
