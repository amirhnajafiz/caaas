package http

// Config for http handler.
type Config struct {
	Port       int  `koanf:"port"`
	EnableAuth byte `koanf:"enable_auth"`
}
