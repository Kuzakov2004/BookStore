package service

import (
	"context"
	"fmt"
	"log"

	"BookStore/internal/publisher"
	"BookStore/internal/publisher/repo"
)

type PublisherService interface {
	GetPublishers(ctx context.Context, page, count int) (lst []*publisher.Publisher, total int, e error)
	GetPublisher(ctx context.Context, id int64) (*publisher.Publisher, error)
}

type publisherService struct {
	repo repo.PublisherRepo
}

func NewPublisherService(r repo.PublisherRepo) (PublisherService, error) {
	return &publisherService{
		repo: r,
	}, nil
}

func (s *publisherService) GetPublishers(ctx context.Context, page, count int) (lst []*publisher.Publisher, total int, e error) {
	lst, e = s.repo.GetPublishers(ctx, page, count)
	if e != nil {
		log.Println("GetPublishers", "Error get publishers ", " [", e, "]")
		return nil, 0, fmt.Errorf("error get publishers [%w]", e)
	}

	total = 0
	total, e = s.repo.GetPublishersCnt(ctx)
	if e != nil {
		log.Println("GetPublishers", "Error get publishers cnt ", " [", e, "]")
		return nil, 0, fmt.Errorf("error get publishers [%w]", e)
	}

	return lst, total, nil
}

func (s *publisherService) GetPublisher(ctx context.Context, id int64) (*publisher.Publisher, error) {
	p, e := s.repo.GetPublisher(ctx, id)
	if e != nil {
		log.Println("GetPublisher", "Error get publisher ", id, " [", e, "]")
		return nil, fmt.Errorf("error get publisher [%w]", e)
	}

	return p, nil
}
