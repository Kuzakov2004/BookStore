package service

import (
	"context"
	"fmt"
	"log"

	"BookStore/internal/warehouse"
	"BookStore/internal/warehouse/repo"
)

type WarehouseService interface {
	GetWarehouses(ctx context.Context, page, count int) (lst []*warehouse.Warehouse, total int, e error)
	GetWarehouse(ctx context.Context, id int64) (*warehouse.Warehouse, error)

	GetWarehouseBooks(ctx context.Context, id, page, count int) (lst []*warehouse.WarehouseBooks, total int, e error)
	GetWarehouseBooksCnt(ctx context.Context, id int) (total int, e error)
}

type warehouseService struct {
	repo repo.WarehousesRepo
}

func NewWarehouseService(r repo.WarehousesRepo) (WarehouseService, error) {
	return &warehouseService{
		repo: r,
	}, nil
}


func (s *warehouseService) GetWarehouses(ctx context.Context, page, count int) (lst []*warehouse.Warehouse, total int, e error) {
	lst, e = s.repo.GetWarehouses(ctx, page, count)
	if e != nil {
		log.Println("GetWarehouses", "Error get Warehouses ", " [", e, "]")
		return nil, 0, fmt.Errorf("error get Warehouses [%w]", e)
	}

	total = 0
	total, e = s.repo.GetWarehousesCnt(ctx)
	if e != nil {
		log.Println("GetWarehouses", "Error get Warehouses cnt ", " [", e, "]")
		return nil, 0, fmt.Errorf("error get Warehouses [%w]", e)
	}

	return lst, total, nil
}

func (s *warehouseService) GetWarehouse(ctx context.Context, id int64) (*warehouse.Warehouse, error) {
	p, e := s.repo.GetWarehouse(ctx, id)
	if e != nil {
		log.Println("GetWarehouse", "Error get Warehouse ", id, " [", e, "]")
		return nil, fmt.Errorf("error get Warehouse [%w]", e)
	}

	return p, nil
}



func (s *warehouseService) GetWarehouseBooks(ctx context.Context, id, page, count int) (lst []*warehouse.WarehouseBooks, total int, e error) {
	lst, e = s.repo.GetWarehouseBooks(ctx, id, page, count)
	if e != nil {
		log.Println("GetWarehouses", "Error get Warehouses ", " [", e, "]")
		return nil, 0, fmt.Errorf("error get Warehouses [%w]", e)
	}

	total = 0
	total, e = s.repo.GetWarehouseBooksCnt(ctx, id)
	if e != nil {
		log.Println("GetWarehouses", "Error get Warehouses cnt ", " [", e, "]")
		return nil, 0, fmt.Errorf("error get WarehousesBooks [%w]", e)
	}

	return lst, total, nil
}

func (s *warehouseService) GetWarehouseBooksCnt(ctx context.Context, id int) (int, error) {
	return s.repo.GetWarehouseBooksCnt(ctx, id)
}

