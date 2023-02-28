package storage

// Config stores storage connection parameters.
type Config struct {
	Host     string `koanf:"host"`
	Port     int    `koanf:"port"`
	User     string `koanf:"user"`
	Pass     string `koanf:"pass"`
	Database string `koanf:"database"`
}
