package model

type UserRole string

const (
	Admin   UserRole = "admin"
	Visitor UserRole = "visitor"
)

func GetRole(role string) UserRole {
	switch role {
	case "admin":
		return Admin
	default:
		return Visitor
	}
}

type CountModel struct {
	TotalData int64 `json:"total_data"`
}
