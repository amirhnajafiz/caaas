package auth

// Config stores data for jwt tokens.
type Config struct {
	PrivateKey string `koanf:"private_key"`
	ExpireTime int    `koanf:"expire_time"`
}
