package http

// Config for http handler.
type Config struct {
	Port       int  `koanf:"port"`
	EnableAuth bool `koanf:"enable_auth"`
}
