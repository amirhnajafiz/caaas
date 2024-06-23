package v1

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserGroupQuery struct {
	Username string `query:"username"`
	Group    string `query:"group"`
}
