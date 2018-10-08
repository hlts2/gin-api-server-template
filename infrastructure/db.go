package infrastructure

import (
	"os"
)

// NewDB returns mysql db object
func NewDB() {
	_ = getENV("DB_USER", "root")
	_ = getENV("DB_PASSWORD", "root")
}

func getENV(key, def string) string {
	env := os.Getenv(key)
	if len(env) == 0 {
		return def
	}

	return env
}
