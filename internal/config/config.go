package config

import (
	"log"
	"strings"

	"github.com/amirhnajafiz/caaas/internal/monitoring/logger"
	"github.com/amirhnajafiz/caaas/internal/monitoring/metrics"
	"github.com/amirhnajafiz/caaas/internal/storage"
	"github.com/amirhnajafiz/caaas/pkg/jwt"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

const (
	// Prefix indicates environment variables prefix.
	Prefix = "CAAAS_"
)

// Config stores the application parameters.
type Config struct {
	Mode           string         `koanf:"mode"`
	HTTPServerPort int            `koanf:"http_server_port"`
	Metrics        metrics.Config `koanf:"metrics"`
	Auth           jwt.Config     `koanf:"auth"`
	Logger         logger.Config  `koanf:"logger"`
	Storage        storage.Config `koanf:"storage"`
}

// LoadConfigs returns the config struct.
func LoadConfigs(path string) Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Printf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider(path), yaml.Parser()); err != nil {
		log.Printf("error loading config.yaml file: %v\n", err)
	}

	// load environment variables
	if err := k.Load(env.Provider(Prefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, Prefix)), "__", ".")
	}), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Printf("error unmarshalling config: %v\n", err)
	}

	return instance
}
