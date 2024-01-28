package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Database DatabaseConfig `toml:"database"`
	Host     string         `toml:"host"`
	Port     uint16         `toml:"port"`
}

type DatabaseConfig struct {
	DatabaseName string `toml:"database_name"`
	Username     string `toml:"username"`
	Password     string `toml:"password"`
	Host         string `toml:"host"`
	Port         uint16 `toml:"port"`
}

// ReadConfig Read config file
func (c *Config) ReadConfig(path string) {
	if _, err := toml.DecodeFile(path, &c); err != nil {
		log.Fatal(err)
	}
}
