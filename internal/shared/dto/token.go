package dto

type UserToken struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	FullName string `json:"full_name"`
}
