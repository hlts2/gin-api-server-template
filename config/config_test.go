package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	tests := []struct {
		path     string
		expected struct {
			hasErr bool
			config *Config
		}
	}{
		{
			path: "./testdata/config.yaml",
			expected: struct {
				hasErr bool
				config *Config
			}{
				hasErr: false,
				config: &Config{
					Server: Server{
						Port: "8080",
					},
					DB: DB{
						User:     "little",
						Password: "tiny",
						Host:     "127.0.0.1",
						Port:     "3306",
						Name:     "sample",
						Param:    "charset=utf8",
					},
				},
			},
		},
		{
			path: "./testdata/bad_config.yaml",
			expected: struct {
				hasErr bool
				config *Config
			}{
				hasErr: true,
				config: nil,
			},
		},
		{
			path: "./testdata/nothing_config.yaml",
			expected: struct {
				hasErr bool
				config *Config
			}{
				hasErr: true,
				config: nil,
			},
		},
	}

	for i, test := range tests {
		func() {
			defer func() {
				config = nil
			}()

			got := LoadConfig(test.path)

			hasErr := !(got == nil)

			if test.expected.hasErr != hasErr {
				t.Errorf("tests[%d] - LoadConfig hasErr is wrong. expected: %v, got: %v", i, test.expected.hasErr, hasErr)
			}

			if !reflect.DeepEqual(test.expected.config, config) {
				t.Errorf("tests[%d] - LoadConfig config is wrong. expected: %v, got: %v", i, test.expected.config, config)
			}
		}()
	}
}

func TestGetConfig(t *testing.T) {
	got := GetConfig()
	expected := Config{}

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("0 - GetConfig config is wrong. expected: %v, got: %v", expected, got)
	}

	config = &Config{DB: DB{User: "little"}}

	got = GetConfig()
	expected = Config{DB: DB{User: "little"}}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("1 - GetConfig config is wrong. expected: %v, got: %v", expected, got)
	}
}
