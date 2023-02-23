package config

import (
	"github.com/creasty/defaults"
)

var (
	C = &Config{}
)

type Config struct {
	LogLevel string `default:"DEBUG" yaml:"logLevel"`
	Mode     string `default:"debug" yaml:"mode"`

	HTTP     HTTP     `yaml:"http"`
	Database Database `yaml:"database"`
}

type HTTP struct {
	Listen   string   `default:":12345" yaml:"listen"`
	Sessions Sessions `yaml:"sessions"`
}

type Sessions struct {
	CookieSecret string `yaml:"cookieSecret"`
}

type Database struct {
	// refer to https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	DSN string `yaml:"dsn"`
}

func init() {
	// Set defaults on start
	defaults.Set(C)
}
