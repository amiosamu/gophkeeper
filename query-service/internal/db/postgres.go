package db

import (
	"log"

	"github.com/amiosamu/gophkeeper/query-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

var _ Storage = (*Postgres)(nil)

func NewPostgres(url string) Storage {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(&models.Record{}); err != nil {
		log.Fatal(err)
	}

	return &Postgres{db}
}

func (p *Postgres) GetRecord(value *models.Record) (*[]models.Record, error) {
	var item = &models.Record{
		UserId: value.UserId,
	}

	if value.MessageType != 0 {
		item.MessageType = value.MessageType
	}

	var result []models.Record

	res := p.db.Model(&models.Record{}).Where(&models.Record{}, item).Scan(&result)
	if res.Error != nil {
		return nil, res.Error
	}

	return &result, nil
}
