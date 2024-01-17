package dto

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type (
	LoginParam struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		Role         string `json:"role"`
		RefreshToken string `json:"refresh_token"`
		ExpiresAt    string `json:"expires_at"`
	}
)

func (l *LoginParam) Validate() error {
	var fieldRules []*validation.FieldRules
	fieldRules = append(fieldRules, validation.Field(&l.Username, is.Email))
	fieldRules = append(fieldRules, validation.Field(&l.Username, validation.Required))
	fieldRules = append(fieldRules, validation.Field(&l.Password, validation.Required, validation.Length(8, 0)))
	return validation.ValidateStruct(l, fieldRules...)
}
