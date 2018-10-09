package infrastructure

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/hlts2/gin-server-template/config"
)

// NewDBConnection returns mysql db connection
func NewDBConnection() (*gorm.DB, error) {
	dsn := getDSN(config.GetConfig().DB)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDSN(c config.DB) string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.Name)
	if len(c.Param) == 0 {
		return dsn
	}
	return dsn + "?" + c.Param
}
