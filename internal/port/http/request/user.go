package request

// Register a new user request.
type Register struct {
	StudentNumber string `json:"student_number"`
	Password      string `json:"password"`
}
