package storage

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewConnection to mysql server.
func NewConnection(cfg Config) (*gorm.DB, error) {
	address := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	return gorm.Open(mysql.Open(address), &gorm.Config{})
}
