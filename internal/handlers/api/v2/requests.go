package v2

type UserRoleQuery struct {
	Username string `query:"username"`
	Role     string `query:"role"`
}
