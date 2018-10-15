package config

import (
	"os"

	yaml "gopkg.in/yaml.v2"
)

var config *Config

// Config is base config object of this application
type Config struct {
	Server `yaml:"server"`
	DB     `yaml:"db"`
}

// Server is config of server
type Server struct {
	Port string `yaml:"port"`
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

// Init loads config of the given path
func Init(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	var c Config
	err = yaml.NewDecoder(f).Decode(&c)
	if err != nil {
		return err
	}
	config = &c

	return nil
}

// GetConfig returns Config object
func GetConfig() Config {
	if config == nil {
		return Config{}
	}

	return *config
}
