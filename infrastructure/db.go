package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/hlts2/gin-server-template/config"
)

// NewDBConnection returns mysql db connection
func NewDBConnection(c *config.DB) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Asia/Tokyo", c.User, c.Password, c.Host, c.Port, c.Name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}
