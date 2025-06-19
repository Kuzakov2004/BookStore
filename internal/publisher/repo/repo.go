package repo

import (
	"BookStore/internal/publisher"
	"context"
)

type PublisherRepo interface {
	GetPublishers(ctx context.Context, page, count int) (lst []*publisher.Publisher, e error)
	GetPublishersCnt(ctx context.Context) (total int, e error)
	GetPublisher(ctx context.Context, id int64) (*publisher.Publisher, error)
}
