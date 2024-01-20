package model

import "time"

type User struct {
	ID        int64     `json:"id" bun:",pk"`
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"password"`
	IsActive  bool      `json:"is_active"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u User) GetRole() UserRole {
	return GetRole(u.Role)
}

type InfoSaving struct {
	ID        int64     `json:"id" bun:",pk"`
	Name      string    `json:"name"`
	Value     int64     `json:"value"`
	UserID    int64     `json:"user_id"`
	User      User      `json:"user" bun:"rel:belongs-to,join:user_id=id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UsedMoney struct {
	ID          int64     `json:"id" bun:",pk"`
	Purpose     string    `json:"purpose"`
	Description string    `json:"description"`
	UserID      int64     `json:"user_id"`
	User        User      `json:"user" bun:"rel:belongs-to,join:user_id=id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type HistoryMoney struct {
	ID        int64     `json:"id" bun:",pk"`
	Value     int64     `json:"value"`
	Operation string    `json:"operation"`
	Purpose   *string   `json:"purpose"`
	UserID    int64     `json:"user_id"`
	User      User      `json:"user" bun:"rel:belongs-to,join:user_id=id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TargetUsedMoney struct {
	ID      int64  `json:"id" bun:",pk"`
	Purpose string `json:"purpose"`
	UserID  int64  `json:"user_id"`
	User    User   `json:"user" bun:"rel:belongs-to,join:user_id=id"`

	Value int64 `json:"value"`
	Month int   `json:"month"`
	Year  int   `json:"year"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
