package service

import (
	"context"
	"elestial/config"
	"elestial/internal/repository"
	"elestial/model"
	"errors"
	"log"
	"time"
)

type AuthService interface {
	Register(ctx context.Context, user model.RegisterRequest) error
	Login(ctx context.Context, user model.User) (model.TokenPair, error)

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
	user.Password = hashedPassword
	if err != nil {
		log.Fatal(err)
	}

	//username check
	if _, err := a.UserRepo.GetUserByName(ctx, user.Name); err == nil {
		return errors.New(" Username exist! ")
	}

	return a.UserRepo.CreateUser(ctx, user)
}

func (a *authService) Login(ctx context.Context, user model.User) (model.TokenPair, error) {
	u, err := a.UserRepo.GetUserByName(ctx, user.Name)
	if err != nil {
		return model.TokenPair{}, err
	}
	if !checkPasswordHash(user.Password, u.Password) {
		return model.TokenPair{}, errors.New(" Wrong password ")
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
