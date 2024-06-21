package gateway

type ClaimResponse struct {
	Username string `json:"username"`
}

type GroupsResponse struct {
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
}
