package jwt

// Config stores data for generating jwt tokens.
type Config struct {
	PrivateKey       string `koanf:"private_key"`
	EncryptionSalt   string `koanf:"encryption_salt"`
	TokensExpireTime int    `koanf:"tokens_expire_time"`
}
