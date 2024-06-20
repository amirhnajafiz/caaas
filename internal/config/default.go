package config

import (
	"github.com/amirhnajafiz/caaas/internal/enum"
	"github.com/amirhnajafiz/caaas/internal/monitoring/logger"
	"github.com/amirhnajafiz/caaas/internal/monitoring/metrics"
	"github.com/amirhnajafiz/caaas/internal/storage"
	"github.com/amirhnajafiz/caaas/pkg/jwt"
)

func Default() Config {
	return Config{
		Mode:           enum.ModeGW,
		HTTPServerPort: 8080,
		Metrics: metrics.Config{
			Enable: true,
			Port:   8081,
		},
		Auth: jwt.Config{
			PrivateKey:       "secret",
			TokensExpireTime: 30, // in miutes
			EncryptionSalt:   "salt",
		},
		Logger: logger.Config{
			Level: "debug",
		},
		Storage: storage.Config{
			Port:     3306,
			Host:     "127.0.0.1",
			User:     "root",
			Pass:     "Amir2222",
			Database: "authx",
		},
	}
}
