package request

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
