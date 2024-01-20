package dto

type (
	GetUsedMoneyParam struct {
		LastID    int64  `json:"last_id"`
		Limit     int    `json:"limit"`
		OrderType int    `json:"order_type"`
		UserID    int64  `json:"user_id"`
		Month     int    `json:"month"`
		Year      int    `json:"year"`
		Day       int    `json:"day"`
		Key       string `json:"key"`
	}
)
