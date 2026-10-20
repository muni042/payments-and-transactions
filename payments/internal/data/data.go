package data

import (
	"payments/internal/conf"

	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(NewData, NewPaymentRepo)

type Data struct {
	DB *gorm.DB
}

func NewData(c *conf.Data) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
