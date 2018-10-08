package config

// Config is base config object of this application
type Config struct {
	DB
}

// DB is config of db
type DB struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Name     string `yaml:"name"`
	Param    string `yaml:"param"`
}

// LoadConfig loads config of the given path
func LoadConfig(path string) error {
	return nil
}
