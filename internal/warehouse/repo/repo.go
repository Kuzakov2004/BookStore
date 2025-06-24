package repo

import (
	"BookStore/internal/warehouse"
	"context"
)

type WarehousesRepo interface {
	GetWarehouses(ctx context.Context, page, count int) (lst []*warehouse.Warehouse, e error)
	GetWarehousesCnt(ctx context.Context) (total int, e error)
	GetWarehouse(ctx context.Context, id int64) (*warehouse.Warehouse, error)

	GetWarehouseBooks(ctx context.Context, id, page, count int) (lst []*warehouse.WarehouseBooks, e error)
	GetWarehouseBooksCnt(ctx context.Context, id int) (total int, e error)
	
}

