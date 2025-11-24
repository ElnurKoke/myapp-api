package service

import (
	"context"
	"elestial/config"
	"elestial/internal/apperror"
	"elestial/internal/repository"
	"elestial/model"
	"errors"
	"fmt"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, user model.RegisterRequest) error
	Login(ctx context.Context, user model.User) (model.TokenPair, error)
	Logout(ctx context.Context, token string) error

	GenerateRefreshToken(userID int) (string, error)
	GenerateAccessToken(userID int) (string, error)
	ParseToken(tokenStr string, secret []byte) (*model.Claims, error)
}

type authService struct {
	AuthRepo repository.AuthRepo
	UserRepo repository.UserRepo
	cfg      *config.Config
}

func NewAuth(AuthRepo repository.AuthRepo, UserRepo repository.UserRepo, cfg *config.Config) *authService {
	return &authService{
		AuthRepo: AuthRepo,
		UserRepo: UserRepo,
		cfg:      cfg,
	}
}

func (a *authService) Register(ctx context.Context, user model.RegisterRequest) error {
	if err := validUser(user); err != nil {
		return err
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return errors.New(" failed to hash password! ")
	}
	user.Password = hashedPassword

	if _, err := a.UserRepo.GetUserByName(ctx, user.Name); err == nil {
		return apperror.ErrUserNameExists
	}

	return a.UserRepo.CreateUser(ctx, user)
}

func (a *authService) Login(ctx context.Context, user model.User) (model.TokenPair, error) {
	u, err := a.UserRepo.GetUserByName(ctx, user.Name)
	if err != nil {
		return model.TokenPair{}, err
	}
	if !checkPasswordHash(user.Password, u.Password) {
		return model.TokenPair{}, apperror.ErrWrongPassword
	}

	access, err := a.GenerateAccessToken(u.ID)
	if err != nil {
		return model.TokenPair{}, err
	}
	refresh, err := a.GenerateRefreshToken(u.ID)
	if err != nil {
		return model.TokenPair{}, err
	}
	expired := time.Now().Add(7 * 24 * time.Hour)
	a.AuthRepo.SaveRefreshToken(ctx, u.ID, refresh, expired)
	return model.TokenPair{Access: access, Refresh: refresh}, nil
}

func (a *authService) Logout(ctx context.Context, refresh string) error {
	_, err := a.ParseToken(refresh, []byte(a.cfg.JWT.RefreshSecret))
	if err != nil {
		return err
	}

	rt, err := a.AuthRepo.GetRefreshToken(ctx, refresh)
	if err != nil {
		return err
	}

	if rt.Revoked {
		return errors.New("refresh token already revoked")
	}

	if err := a.AuthRepo.RevokeRefreshToken(ctx, rt.ID); err != nil {
		return fmt.Errorf("failed to revoke refresh token: %w", err)
	}

	return nil
}
