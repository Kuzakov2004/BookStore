package repo

import (
	"BookStore/internal/book"
	"BookStore/internal/order"
	"context"
)

type OrderRepo interface {
	GetOrders(ctx context.Context, status string, page, count int) (lst []*order.Order, e error)
	GetOrdersCnt(ctx context.Context, status string) (total int, e error)
	GetOrder(ctx context.Context, id int64) (*order.OrderDetail, error)
	CreateOrder(ctx context.Context) (int64, error)
	SetOrderClient(ctx context.Context, orderId int64, clientId int64) error

	SaveShip(ctx context.Context, orderId int64, ship *order.Ship) error
	AddBooks(ctx context.Context, orderId int64, books []int64) error
	SaveBookQty(ctx context.Context, orderId int64, books []int64, qty []int) error
	Pay(ctx context.Context, orderId int64, books []int64, qty []int) error
	Send(ctx context.Context, orderId int64) error
	DelBookFromOrder(ctx context.Context, orderId int64, bookId int64) error

	FindClient(ctx context.Context, str string) (lst []*order.Client, e error)
	FindBook(ctx context.Context, orderId int64, page, count int) (lst []*book.Book, e error)
}
