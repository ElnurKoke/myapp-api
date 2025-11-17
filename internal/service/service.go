package service

import (
	"elestial/config"
	"elestial/internal/repository"
)

type Service struct {
	AuthService
}

func NewService(repository *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		AuthService: NewAuth(repository.AuthRepo, repository.UserRepo, cfg),
	}
}
