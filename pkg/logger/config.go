package logger

// Config parameters for zap logger.
type Config struct {
	Level string `koanf:"level"`
}
