package service

import (
	"BookStore/internal/admin/repo"
	"BookStore/internal/book"
	"context"
	"fmt"
	"log"
)

type AdminService interface {
	UpdateBook(ctx context.Context, info *book.FullInfo) error
	CreateBook(ctx context.Context, info *book.FullInfo) (int64, error)
	DeleteBook(ctx context.Context, id int64) error
}

type adminService struct {
	repo repo.AdminRepo
}

func NewAdminService(r repo.AdminRepo) (AdminService, error) {
	return &adminService{
		repo: r,
	}, nil
}

func (s *adminService) UpdateBook(ctx context.Context, info *book.FullInfo) error {
	e := s.repo.UpdateBook(ctx, info)
	if e != nil {
		log.Println("Login", "Error update book", info, " [", e, "]")
		return fmt.Errorf("error update book [%w]", e)
	}

	return nil
}

func (s *adminService) CreateBook(ctx context.Context, info *book.FullInfo) (int64, error) {
	id, e := s.repo.CreateBook(ctx, info)
	if e != nil {
		log.Println("Login", "Error update book", info, " [", e, "]")
		return 0, fmt.Errorf("error update book [%w]", e)
	}

	return id, nil
}

func (s *adminService) DeleteBook(ctx context.Context, id int64) error {
	e := s.repo.DeleteBook(ctx, id)
	if e != nil {
		log.Println("Error delete book", id, " [", e, "]")
		return fmt.Errorf("error delete book [%w]", e)
	}

	return nil
}
