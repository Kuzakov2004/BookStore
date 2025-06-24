package service

import (
	"BookStore/internal/order"
	repo2 "BookStore/internal/order/repo"
	"context"
	"fmt"
	"log"
)

type OrderService interface {
	GetOrders(ctx context.Context, status string, page, count int) (lst []*order.Order, total int, e error)
	GetOrder(ctx context.Context, id int64) (*order.OrderDetail, error)
	SetOrderClient(ctx context.Context, orderID int64, clientId int64) error
	CreateOrder(ctx context.Context) (int64, error)

	FindClient(ctx context.Context, str string) (lst []*order.Client, e error)
}

type orderService struct {
	repo repo2.OrderRepo
}

func NewOrderService(r repo2.OrderRepo) (OrderService, error) {
	return &orderService{
		repo: r,
	}, nil
}

func (s *orderService) GetOrders(ctx context.Context, status string, page, count int) (lst []*order.Order, total int, e error) {
	lst, e = s.repo.GetOrders(ctx, status, page, count)
	if e != nil {
		log.Println("GetOrders", "Error get orders ", status, " [", e, "]")
		return nil, 0, fmt.Errorf("error get orders [%w]", e)
	}

	total = 0
	total, e = s.repo.GetOrdersCnt(ctx, status)
	if e != nil {
		log.Println("GetBooks", "Error get orders cnt ", status, " [", e, "]")
		return nil, 0, fmt.Errorf("error get orders [%w]", e)
	}

	return lst, total, nil
}

func (s *orderService) GetOrder(ctx context.Context, id int64) (*order.OrderDetail, error) {
	o, e := s.repo.GetOrder(ctx, id)
	if e != nil {
		log.Println("GetOrder", "Error get order ", id, " [", e, "]")
		return nil, fmt.Errorf("error get order [%w]", e)
	}

	return o, nil
}

func (s *orderService) FindClient(ctx context.Context, str string) (lst []*order.Client, e error) {
	lst, e = s.repo.FindClient(ctx, str)
	if e != nil {
		log.Println("FindClient", "Error find clients", str, " [", e, "]")
		return nil, fmt.Errorf("error find clients [%w]", e)
	}

	return lst, nil
}

func (s *orderService) CreateOrder(ctx context.Context) (int64, error) {
	id, e := s.repo.CreateOrder(ctx)
	if e != nil {
		log.Println("CreateOrder", "Error create order", " [", e, "]")
		return 0, fmt.Errorf("error create order [%w]", e)
	}

	return id, nil
}

func (s *orderService) SetOrderClient(ctx context.Context, orderID int64, clientId int64) error {
	return nil
}
