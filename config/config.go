package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

var config *Config

// Config is base config object of this application
type Config struct {
	DB `yaml:"db"`
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
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	err = yaml.NewDecoder(f).Decode(config)
	if err != nil {
		return err
	}

	return nil
}

// GetConfig returns Config object
func GetConfig() Config {
	return *config
}
