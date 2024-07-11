package auth

import (
	"context"
	"onlineShop/infra/response"
	"onlineShop/internal/config"
)

type Repository interface {
	CreateAuth(ctx context.Context, model AuthEntity) (err error)
	GetAuthByEmail(ctx context.Context, email string) (model AuthEntity, err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) register(ctx context.Context, req RegisterRequestPayload) (err error) {
	authEntity := NewFromRegisterRequest(req)

	err = authEntity.Validate()
	if err != nil {
		return
	}

	err = authEntity.EncryptPassword(int(config.Cfg.App.Encryption.Salt))
	if err != nil {
		return
	}

	model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		if err != response.ErrNotFound {
			return
		}
	}

	if model.IsExists() {
		return response.ErrEmailAlreadyExist
	}

	return s.repo.CreateAuth(ctx, authEntity)
}

func (s service) login(ctx context.Context, req LoginRequestPayload) (token string, err error) {
	authEntity := NewFromLoginRequest(req)

	err = authEntity.ValidateEmail()
	if err != nil {
		return
	}

	err = authEntity.ValidatePassword()
	if err != nil {
		return
	}

	model, err := s.repo.GetAuthByEmail(ctx, authEntity.Email)
	if err != nil {
		return
	}

	err = authEntity.VerifyPasswordFromPlain(model.Password)
	if err != nil {
		err = response.ErrPasswordNotMatch
		return
	}

	token, err = model.GenerateToken(config.Cfg.App.Encryption.JWTSecret)
	return
}
