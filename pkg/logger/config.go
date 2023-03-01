package logger

// Config parameters for logger.
type Config struct {
	Level  string `koanf:"level"`
	Enable bool   `koanf:"enable"`
}
