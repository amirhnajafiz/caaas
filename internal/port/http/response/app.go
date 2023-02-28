package response

// AppResponse data for a single app.
type AppResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Clients []uint `json:"clients"`
}
