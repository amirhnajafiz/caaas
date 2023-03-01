package config

import (
	"github.com/amirhnajafiz/authX/internal/port/http"
	"github.com/amirhnajafiz/authX/internal/storage"
	"github.com/amirhnajafiz/authX/pkg/auth"
	"github.com/amirhnajafiz/authX/pkg/logger"
)

func Default() Config {
	return Config{
		Auth: auth.Config{
			PrivateKey: "private",
			ExpireTime: 5,
		},
		Logger: logger.Config{
			Level:  "debug",
			Enable: false,
		},
		HTTP: http.Config{
			Port: 5000,
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
