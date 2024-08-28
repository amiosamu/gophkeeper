package db

import "github.com/amiosamu/gophkeeper/auth-service/internal/models"

type Storage interface {
	AddUser(user *models.User) (*models.User, error)
	GetUser(user *models.User) (*models.User, error)
}
