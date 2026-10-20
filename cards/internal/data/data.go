package data

import (
	//"cards/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"cards/internal/conf"
)

// ProviderSet is data providers.
var DataProviderSet = wire.NewSet(NewData, NewDB, NewCardRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewDB(c *conf.Data) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(c.Database.Source), &gorm.Config{})
}
