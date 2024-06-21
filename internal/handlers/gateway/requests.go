package gateway

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
