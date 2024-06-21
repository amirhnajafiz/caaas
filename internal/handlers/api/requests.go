package api

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserGroupRequest struct {
	Username string   `json:"username"`
	Groups   []string `json:"groups"`
}
