package service

import (
	"BookStore/internal/admin/repo"
	"BookStore/internal/book"
	"BookStore/internal/publisher"
	"context"
	"fmt"
	"log"
)

type AdminService interface {
	UpdateBook(ctx context.Context, info *book.FullInfo) error
	CreateBook(ctx context.Context, info *book.FullInfo) (int64, error)
	DeleteBook(ctx context.Context, id int64) error

	UpdatePublisher(ctx context.Context, p *publisher.Publisher) error
	CreatePublisher(ctx context.Context, p *publisher.Publisher) (int64, error)
	DeletePublisher(ctx context.Context, id int64) error
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
		log.Println("UpdateBook", "Error update book", info, " [", e, "]")
		return fmt.Errorf("error update book [%w]", e)
	}

	return nil
}

func (s *adminService) CreateBook(ctx context.Context, info *book.FullInfo) (int64, error) {
	id, e := s.repo.CreateBook(ctx, info)
	if e != nil {
		log.Println("CreateBook", "Error update book", info, " [", e, "]")
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

func (s *adminService) UpdatePublisher(ctx context.Context, p *publisher.Publisher) error {
	e := s.repo.UpdatePublisher(ctx, p)
	if e != nil {
		log.Println("UpdatePublisher", "Error update publisher", p, " [", e, "]")
		return fmt.Errorf("error update publisher [%w]", e)
	}

	return nil
}
func (s *adminService) CreatePublisher(ctx context.Context, p *publisher.Publisher) (int64, error) {
	id, e := s.repo.CreatePublisher(ctx, p)
	if e != nil {
		log.Println("CreatePublisher", "Error update publisher", p, " [", e, "]")
		return 0, fmt.Errorf("error update publisher [%w]", e)
	}

	return id, nil
}
func (s *adminService) DeletePublisher(ctx context.Context, id int64) error {
	e := s.repo.DeletePublisher(ctx, id)
	if e != nil {
		log.Println("Error delete publisher", id, " [", e, "]")
		return fmt.Errorf("error delete publisher [%w]", e)
	}

	return nil
}
