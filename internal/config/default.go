package config

import (
	"github.com/amirhnajafiz/caaas/internal/storage"
	"github.com/amirhnajafiz/caaas/pkg/auth"
	"github.com/amirhnajafiz/caaas/pkg/logger"
)

func Default() Config {
	return Config{
		Auth: auth.Config{
			PrivateKey: "private",
			ExpireTime: 5,
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
