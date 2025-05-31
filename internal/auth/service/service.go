package service

import (
	"context"
	"fmt"
	"log"

	"BookStore/internal/auth/repo"
)

type AuthAdminService interface {
	Login(ctx context.Context, user, pass string) error
	Logout(ctx context.Context) error
}

type authAdminService struct {
	repo repo.AdminRepo
}

func NewAdminService(r repo.AdminRepo) (AuthAdminService, error) {
	return &authAdminService{
		repo: r,
	}, nil
}

func (s *authAdminService) Login(ctx context.Context, user, pass string) (int64, error) {
	id, e := s.repo.Login(ctx, user, pass)
	if e != nil {
		log.Println("Login", "Error login", user, " [", e, "]")
		return id, fmt.Errorf("error login [%w]", e)
	}

	return id, nil
}

func (s *authAdminService) Logout(ctx context.Context) error {
	return s.repo.Logout(ctx)
}
