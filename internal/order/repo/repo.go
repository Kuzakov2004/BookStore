package repo

import (
	"BookStore/internal/order"
	"context"
)

type OrderRepo interface {
	GetOrders(ctx context.Context, status string, page, count int) (lst []*order.Order, e error)
	GetOrdersCnt(ctx context.Context, status string) (total int, e error)
	GetOrder(ctx context.Context, id int64) (*order.OrderDetail, error)
	CreateOrder(ctx context.Context) (int64, error)

	FindClient(ctx context.Context, str string) (lst []*order.Client, e error)
}
