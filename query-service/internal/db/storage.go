package db

import "github.com/amiosamu/gophkeeper/query-service/internal/models"

type Storage interface {
	GetRecord(value *models.Record) (*[]models.Record, error)
}
