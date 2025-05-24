package book

import (
	"context"
	"fmt"
	"log"

	"BookStore/internal/book/repo"
)

type BookService interface {
	GetBooks(ctx context.Context, genre string, page, count int) (lst []*book.Book, total int, e error)
	FindBooks(ctx context.Context, searchStr string) (lst []*book.Book, e error)
}

type bookService struct {
	repo repo.BookRepo
}

func NewBookService(r repo.BookRepo) BookService {
	return &bookService{
		repo: r,
	}
}

func (s *bookService) GetBooks(ctx context.Context, genre string, page, count int) (lst []*book.Book, total int, e error) {
	lst, e = s.repo.GetBooks(ctx, genre, page, count)
	if e != nil {
		log.Println("GetBooks", "Error get books ", genre, " [", e, "]")
		return nil, 0, fmt.Errorf("error get books [%w]", e)
	}

	total = 0
	total, e = s.repo.GetBooksCnt(ctx, genre)
	if e != nil {
		log.Println("GetBooks", "Error get books cnt ", genre, " [", e, "]")
		return nil, 0, fmt.Errorf("error get books [%w]", e)
	}

	return lst, total, nil
}

func (s *bookService) FindBooks(ctx context.Context, searchStr string) (lst []*book.Book, e error) {
	lst, e = s.repo.Find(ctx, searchStr)
	if e != nil {
		log.Println("FindBooks", "Error find books ", searchStr, " [", e, "]")
		return nil, fmt.Errorf("error find books [%w]", e)
	}

	return lst, nil
}

func (s *bookService) GetBook(ctx context.Context, id int64) (*book.FullInfo, error) {
	b, e := s.repo.GetBook(ctx, id)
	if e != nil {
		log.Println("FindBooks", "Error get book ", id, " [", e, "]")
		return nil, fmt.Errorf("error get books [%w]", e)
	}

	return b, nil
}
