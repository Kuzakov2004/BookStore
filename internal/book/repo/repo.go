package repo

import (
	"context"

	"BookStore/internal/book"
)

type BookRepo interface {
	GetBooks(ctx context.Context, genre string, page, count int) (lst []*book.Book, e error)
	GetBooksCnt(ctx context.Context, genre string) (total int, e error)
	Find(ctx context.Context, findStr string) (lst []*book.Book, e error)
	GetBook(ctx context.Context, id int64) (*book.FullInfo, error)
	GetAuthors(ctx context.Context) (lst []*book.Author, e error)
}
