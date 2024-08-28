package db

import (
	"log"

	"github.com/amiosamu/gophkeeper/command-consumer-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB *gorm.DB
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

func (p *Postgres) AddRecord(value *models.Record) error {
	tx := p.DB.Create(&value)
	return tx.Error
}

func (p *Postgres) ModifyRecord(value *models.Record) error {
	tx := p.DB.Model(&models.Record{Id: value.Id}).
		Updates(models.Record{
			UserId:      value.UserId,
			MessageType: value.MessageType,
			Data:        value.Data,
			Meta:        value.Meta,
		})
	return tx.Error
}

func (p *Postgres) DeleteRecord(value *models.Record) error {
	tx := p.DB.Delete(&models.Record{}, value.Id)
	return tx.Error
}
