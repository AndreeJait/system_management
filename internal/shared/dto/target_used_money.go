package dto

type (
	GetTargetUsedMoneyByIDParam struct {
		LastID    int64 `json:"last_id"`
		Limit     int   `json:"limit"`
		OrderType int   `json:"order_type"`
		UserID    int64 `json:"user_id"`
		Year      int   `json:"year"`
		Month     int   `json:"month"`
	}
)
