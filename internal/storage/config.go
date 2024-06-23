package storage

import "fmt"

// Config stores storage connection parameters.
type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Pass     string `koanf:"pass"`
	Database string `koanf:"database"`
	SSL      bool   `koanf:"ssl"`
}

func (c Config) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%t", c.User, c.Pass, c.Host, c.Port, c.Database, c.SSL)
}
