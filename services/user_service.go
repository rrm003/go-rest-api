package services

import "go-rest-api/models"

type UserService interface {
	SignUp(user models.User) error
	Login(username, password string) (string, error)
	GetUsers() ([]models.User, error)
	GetUser(id string) (models.User, error)
	UpdateUser(id string, user models.User) error
	DeleteUser(id string) error
	GetCountires() ([]string, error)
}
