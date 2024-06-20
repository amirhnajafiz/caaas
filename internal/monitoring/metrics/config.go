package metrics

// Config for metrics server.
type Config struct {
	Enable bool `koanf:"enable"`
	Port   int  `koanf:"port"`
}
