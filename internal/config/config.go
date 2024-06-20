package config

import (
	"log"

	"github.com/amirhnajafiz/caaas/internal/storage"
	"github.com/amirhnajafiz/caaas/pkg/auth"
	"github.com/amirhnajafiz/caaas/pkg/logger"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

// Config stores the application parameters.
type Config struct {
	Auth    auth.Config    `koanf:"auth"`
	Logger  logger.Config  `koanf:"logger"`
	Storage storage.Config `koanf:"storage"`
}

// LoadConfigs returns the config struct.
func LoadConfigs() Config {
	var instance Config

	k := koanf.New(".")

	// load default
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Printf("error loading deafult: %v\n", err)
	}

	// load configs file
	if err := k.Load(file.Provider("config.yaml"), yaml.Parser()); err != nil {
		log.Printf("error loading config.yaml file: %v\n", err)
	}

	// unmarshalling
	if err := k.Unmarshal("", &instance); err != nil {
		log.Printf("error unmarshalling config: %v\n", err)
	}

	return instance
}
