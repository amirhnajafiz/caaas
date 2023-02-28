package response

type AppResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Clients []uint `json:"clients"`
}
