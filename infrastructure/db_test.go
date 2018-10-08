package infrastructure

import (
	"fmt"
	"testing"

	"github.com/hlts2/gin-server-template/config"
)

func TestGetDSN(t *testing.T) {
	var (
		user     = "little"
		password = "tiny"
		host     = "127.0.0.1"
		port     = "3306"
		name     = "sample"
		param    = "charset=utf8mb4&param=true&loc=Asia/Tokyo"
	)

	got := getDSN(&config.DB{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     name,
		Param:    param,
	})

	expected := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, name, param)

	if expected != got {
		t.Errorf("getDSN is wrong. expected: %s, got: %s", expected, got)
	}
}
