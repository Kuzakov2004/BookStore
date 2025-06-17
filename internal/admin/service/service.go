package service

import (
	"BookStore/internal/book"
	"context"
	"fmt"
	"log"

	"BookStore/internal/auth/repo"
)

type AdminService interface {
	UpdateBook(ctx context.Context, info *book.FullInfo) error
	CreateBook(ctx context.Context, info *book.FullInfo) error
}

type adminService struct {
	repo repo.AdminRepo
}

func NewAdminService(r repo.AuthRepo) (AdminService, error) {
	return &adminService{
		repo: r,
	}, nil
}

func (s *adminService) UpdateBook(ctx context.Context, info *book.FullInfo) error {
	e := repo.UpdateBook(ctx, info)
	if e != nil {
		log.Println("Login", "Error login", user, " [", e, "]")
		return id, fmt.Errorf("error login [%w]", e)
	}

	return id, nil
}

func (s *adminService) CreateBook(ctx context.Context, info *book.FullInfo) error {
	return s.repo.Logout(ctx)
}
