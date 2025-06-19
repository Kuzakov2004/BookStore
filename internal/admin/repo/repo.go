package repo

import (
	"BookStore/internal/book"
	"BookStore/internal/publisher"
	"context"
)

type AdminRepo interface {
	UpdateBook(ctx context.Context, info *book.FullInfo) error
	CreateBook(ctx context.Context, info *book.FullInfo) (int64, error)
	DeleteBook(ctx context.Context, id int64) error

	UpdatePublisher(ctx context.Context, p *publisher.Publisher) error
	CreatePublisher(ctx context.Context, p *publisher.Publisher) (int64, error)
	DeletePublisher(ctx context.Context, id int64) error
}
