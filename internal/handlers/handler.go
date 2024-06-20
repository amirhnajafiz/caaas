package handlers

// Handler is a abstract struct of api, gateway, and migrate handers,
// which will be used based on the give mode.
type Handler interface {
	Execute() error
}
