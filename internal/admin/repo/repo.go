package repo

import (
	"BookStore/internal/book"
	"context"
)

type AdminRepo interface {
	UpdateBook(ctx context.Context, info *book.FullInfo) error
	CreateBook(ctx context.Context, info *book.FullInfo) (int64, error)
}
