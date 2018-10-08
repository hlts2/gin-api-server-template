package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/hlts2/gin-server-template/config"
)

// NewDBConnection returns mysql db connection
func NewDBConnection(c *config.DB) (*sql.DB, error) {
	db, err := sql.Open("mysql", getDSN(c))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getDSN(c *config.DB) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", c.User, c.Password, c.Host, c.Port, c.Name, c.Param)
}
