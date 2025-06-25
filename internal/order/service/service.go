package service

import (
	"BookStore/internal/book"
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
	SaveShip(ctx context.Context, orderId int64, ship *order.Ship) error
	AddBooks(ctx context.Context, orderId int64, books []int64) error
	SaveBookQty(ctx context.Context, orderId int64, books []int64, qty []int) error
	Pay(ctx context.Context, orderId int64, books []int64, qty []int) error
	Send(ctx context.Context, orderId int64) error

	FindClient(ctx context.Context, str string) (lst []*order.Client, e error)
	FindBook(ctx context.Context, orderId int64, page, count int) (lst []*book.Book, e error)
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

func (s *orderService) FindBook(ctx context.Context, orderId int64, page, count int) (lst []*book.Book, e error) {
	lst, e = s.repo.FindBook(ctx, orderId, page, count)
	if e != nil {
		log.Println("FindBook", "Error find books", " [", e, "]")
		return nil, fmt.Errorf("error find books [%w]", e)
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

func (s *orderService) SetOrderClient(ctx context.Context, orderId int64, clientId int64) error {
	e := s.repo.SetOrderClient(ctx, orderId, clientId)
	if e != nil {
		log.Println("CreateOrder", "Error create order", " [", e, "]")
		return fmt.Errorf("error create order [%w]", e)
	}

	return nil
}

func (s *orderService) SaveShip(ctx context.Context, orderId int64, ship *order.Ship) error {
	e := s.repo.SaveShip(ctx, orderId, ship)
	if e != nil {
		log.Println("SaveShip", "Error save ship", " [", e, "]")
		return fmt.Errorf("error save ship [%w]", e)
	}

	return nil
}

func (s *orderService) AddBooks(ctx context.Context, orderId int64, books []int64) error {
	e := s.repo.AddBooks(ctx, orderId, books)
	if e != nil {
		log.Println("AddBooks", "Error add books", " [", e, "]")
		return fmt.Errorf("error add books [%w]", e)
	}

	return nil
}

func (s *orderService) SaveBookQty(ctx context.Context, orderId int64, books []int64, qty []int) error {
	e := s.repo.SaveBookQty(ctx, orderId, books, qty)
	if e != nil {
		log.Println("SaveBookQty", "Error save book qty", " [", e, "]")
		return fmt.Errorf("error save book qty [%w]", e)
	}

	return nil
}

func (s *orderService) Pay(ctx context.Context, orderId int64, books []int64, qty []int) error {
	e := s.repo.Pay(ctx, orderId, books, qty)
	if e != nil {
		log.Println("SaveBookQty", "Error pay", " [", e, "]")
		return fmt.Errorf("error pay [%w]", e)
	}

	return nil
}

func (s *orderService) Send(ctx context.Context, orderId int64) error {
	e := s.repo.Send(ctx, orderId)
	if e != nil {
		log.Println("Send", "Error send", " [", e, "]")
		return fmt.Errorf("error send [%w]", e)
	}

	return nil
}
