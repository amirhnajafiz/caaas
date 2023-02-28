package config

type Config struct {
}

func LoadConfigs() Config {
	return Default()
}
