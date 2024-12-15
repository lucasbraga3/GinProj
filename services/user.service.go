package services

import "goproject.com/simple-api/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUserbyName(*string) (*models.User, error)
	GetAll() ([]models.User, error)
	UpdateUser(*string, *models.User) error
	DeleteUser(*string) error
}