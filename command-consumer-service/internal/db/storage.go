package db

import "github.com/amiosamu/gophkeeper/command-consumer-service/internal/models"

type Storage interface {
	AddRecord(value *models.Record) error
	ModifyRecord(value *models.Record) error
	DeleteRecord(value *models.Record) error
}
