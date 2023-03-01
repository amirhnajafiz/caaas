package config

import (
	"github.com/amirhnajafiz/authX/internal/port/http"
	"github.com/amirhnajafiz/authX/internal/storage"
	"github.com/amirhnajafiz/authX/pkg/auth"
)

func Default() Config {
	return Config{
		Auth: auth.Config{
			PrivateKey: "",
			ExpireTime: 5,
		},
		HTTP: http.Config{
			Port:       5000,
			EnableAuth: false,
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
