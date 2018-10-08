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
		param    = "charset=utf8mb4&parseTime=true&loc=Asia/Tokyo"
	)

	tests := []struct {
		c        *config.DB
		expected string
	}{
		{
			c: &config.DB{
				User:     user,
				Password: password,
				Host:     host,
				Port:     port,
				Name:     name,
			},
			expected: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name),
		},
		{
			c: &config.DB{
				User:     user,
				Password: password,
				Host:     host,
				Port:     port,
				Name:     name,
				Param:    param,
			},
			expected: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", user, password, host, port, name, param),
		},
	}

	for i, test := range tests {
		got := getDSN(test.c)

		if test.expected != got {
			t.Errorf("tests[%d] - getDSN is wrong. expected: %v, got: %v", i, test.expected, got)
		}
	}
}
