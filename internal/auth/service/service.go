package service

import (
	"context"
	"fmt"
	"log"

	"BookStore/internal/auth/repo"
)

type AuthService interface {
	Login(ctx context.Context, user, pass string) (int64, error)
	Logout(ctx context.Context) error

	ClientLogin(ctx context.Context, user, pass string) (int64, error)
	ClientLogout(ctx context.Context) error
}

type authService struct {
	repo repo.AuthRepo
}

func NewAuthService(r repo.AuthRepo) (AuthService, error) {
	return &authService{
		repo: r,
	}, nil
}

func (s *authService) Login(ctx context.Context, user, pass string) (int64, error) {
	id, e := s.repo.Login(ctx, user, pass)
	if e != nil {
		log.Println("Login", "Error admin login", user, " [", e, "]")
		return id, fmt.Errorf("error login [%w]", e)
	}

	return id, nil
}

func (s *authService) ClientLogin(ctx context.Context, user, pass string) (int64, error) {
	id, e := s.repo.Login(ctx, user, pass)
	if e != nil {
		log.Println("Login", "Error login", user, " [", e, "]")
		return id, fmt.Errorf("error login [%w]", e)
	}

	return id, nil
}

func (s *authService) Logout(ctx context.Context) error {
	return s.repo.Logout(ctx)
}

func (s *authService) ClientLogout(ctx context.Context) error {
	return s.repo.ClientLogout(ctx)
}
