package impl

import (
	"context"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"system_management/commons/ierr"
	"system_management/commons/password"
	"system_management/commons/utils"
	"system_management/internal/model"
	"system_management/internal/shared/constant"
	"system_management/internal/shared/dto"
	"time"
)

func (u useCase) authenticate(ctx context.Context, username, plainPassword string) (user model.User, err error) {
	repoUser := u.repoManagementSystem.GetUserRepository()

	user, err = repoUser.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, ierr.ErrUserNotFound) {
			return user, ierr.ErrInvalidCred
		}
		return
	}

	if password.ComparePasswords(user.Password, []byte(plainPassword)) {
		if !user.IsActive {
			return user, ierr.ErrUserNotActive
		}
		return user, nil
	}

	return user, ierr.ErrInvalidCred
}

func (u useCase) Login(ctx context.Context, param dto.LoginParam) (dto.LoginResponse, error) {
	var res dto.LoginResponse

	err := param.Validate()
	if err != nil {
		return res, err
	}

	user, err := u.authenticate(ctx, param.Username, param.Password)
	if err != nil {
		return res, err
	}

	accessToken, expiredAt, refreshToken, err := u.generateJWT(ctx, user)
	if err != nil {
		return res, err
	}

	res.AccessToken = accessToken
	res.RefreshToken = refreshToken
	res.ExpiresAt = expiredAt.Format(time.RFC3339)
	res.Role = user.Role
	return res, nil
}

func (u useCase) generateJWT(ctx context.Context, user model.User) (accessToken string, expiresAt time.Time, refreshToken string, err error) {

	//generate access token
	accessToken, expiresAt, err = u.generateAccessToken(ctx, user)
	if err != nil {
		return
	}

	// generate refresh token
	refreshToken, err = u.generateRefreshToken(ctx, user)
	if err != nil {
		return
	}

	return
}

func (u useCase) generateAccessToken(ctx context.Context, user model.User) (accessToken string, expiresAt time.Time, err error) {

	fullNameEncrypted, err := utils.Encrypt(user.FullName, u.cfg.Jwt.EncryptionKey)
	if err != nil {
		return
	}

	emailEncrypted, err := utils.Encrypt(user.Username, u.cfg.Jwt.EncryptionKey)
	if err != nil {
		return
	}
	expiresAt = time.Now().Add(time.Duration(u.cfg.Jwt.TokenExpiration) * time.Minute)
	expiresAtUnix := time.Now().Add(time.Duration(u.cfg.Jwt.TokenExpiration) * time.Minute).Unix()
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"full_name":  fullNameEncrypted,
		"email":      emailEncrypted,
		"exp":        expiresAtUnix,
		"token_type": constant.TokenTypeAccess,
	}).SignedString([]byte(u.cfg.Jwt.SigningKey))
	err = errors.Wrap(err, "cannot generate token")
	return
}

func (u useCase) generateRefreshToken(ctx context.Context, user model.User) (refreshToken string, err error) {

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         user.ID,
		"exp":        time.Now().AddDate(0, 0, u.cfg.Jwt.RefreshTokenExpiration).Unix(),
		"token_type": constant.TokenTypeRefresh,
	}).SignedString([]byte(u.cfg.Jwt.SigningKey))
	err = errors.Wrap(err, "cannot generate token")
	return
}
