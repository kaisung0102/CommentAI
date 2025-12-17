package data

import (
	"review-service/internal/conf"
	"review-service/internal/data/query"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewReviewRepo, NewDB)

// Data .
type Data struct {
	// TODO wrapped database client
	query *query.Query
}

// NewData .
func NewData(db *gorm.DB) (*Data, func(), error) {
	cleanup := func() {
		log.Info("closing the data resources")
	}
	query.SetDefault(db)
	return &Data{query: query.Q}, cleanup, nil
}

func NewDB(c *conf.Data) (*gorm.DB, error) {
	switch strings.ToLower(c.Database.GetDriver()) {
	case "mysql":
		return gorm.Open(mysql.Open(c.Database.GetSource()), &gorm.Config{})
	default:
		return nil, errors.BadRequest("database.driver", "unsupported driver")
	}
}
