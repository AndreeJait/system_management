package dto

type (
	LoginParam struct {
		Username string
		Password string
	}

	LoginResponse struct {
		AccessToken  string
		RefreshToken string
		TimeExpired  int64
	}
)
