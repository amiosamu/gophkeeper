package db

import (
	"errors"
	"log"

	"github.com/amiosamu/gophkeeper/auth-service/internal/models"
	"github.com/amiosamu/gophkeeper/auth-service/internal/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

var _ (Storage) = (*Postgres)(nil)

var (
	ErrorUserNotFound     = errors.New("user not found")
	ErrorUserAlreadyExist = errors.New("user already exist")
)

func NewPostgres(url string) Storage {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal(err)
	}

	return &Postgres{db}
}

func (p *Postgres) GetUser(value *models.User) (*models.User, error) {
	user := models.User{}

	tx := p.db.Where(&models.User{Name: value.Name}).First(&user)

	if tx.Error != nil {
		return &models.User{
			Name: value.Name,
		}, ErrorUserNotFound
	}

	return &user, nil
}

func (p *Postgres) AddUser(value *models.User) (*models.User, error) {
	user, err := p.GetUser(value)

	if err == nil {
		return user, ErrorUserAlreadyExist
	}

	user.Password, err = utils.HashPassword(value.Password)

	if err != nil {
		return nil, err
	}

	tx := p.db.Create(&user)

	return user, tx.Error
}
