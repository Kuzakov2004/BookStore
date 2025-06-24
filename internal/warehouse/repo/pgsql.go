package repo

import (
	"BookStore/internal/warehouse"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type warehouseRepo struct {
	connPool *sql.DB

	queryWarehousesStmt      *sql.Stmt
	getWarehousesCntStmt     *sql.Stmt
	getWarehouseStmt         *sql.Stmt

	queryWarehouseBooksStmt  *sql.Stmt
	getWarehouseBooksCntStmt *sql.Stmt
}

func NewWarehouseRepo(db *sql.DB) (WarehousesRepo, error) {

	r := &warehouseRepo{connPool: db}
	var e error

	r.queryWarehousesStmt, e = db.Prepare(`SELECT w.id, w.address, w.capacity 
											FROM store.warehouses w
    										LIMIT $1 OFFSET $2`)
	if e != nil {
		return nil, e
	}

	r.getWarehousesCntStmt, e = db.Prepare(`SELECT count(*) FROM store.warehouses`)
	if e != nil {
		return nil, e
	}

	r.getWarehouseStmt, e = db.Prepare(`SELECT w.id, w.address, w.capacity 
											FROM store.warehouses w 
											WHERE w.id = $1 
    										LIMIT 1`)
	if e != nil {
		return nil, e
	}

	r.getWarehouseBooksCntStmt, e = db.Prepare(`SELECT COUNT(*) AS book_count FROM store.warehouse_books WHERE wrhs_id = $1;`)

	r.queryWarehouseBooksStmt, e = db.Prepare(`select
    											b.isbn,
   											    b.title,
    											a.first_name || ' ' || a.last_name AS author,
    											b.genre,
    											b.price,
    											p.name AS publisher,
    											wb.qty AS quantity_on_stock
												FROM store.warehouse_books wb
												JOIN store.books b ON wb.book_id = b.id
												JOIN store.authors a ON b.author_id = a.id
												JOIN store.publishers p ON b.publisher_id = p.id
												WHERE wb.wrhs_id = $1
												ORDER BY b.title
												LIMIT $2 OFFSET $3`)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (r *warehouseRepo) GetWarehouses(ctx context.Context, page, count int) (lst []*warehouse.Warehouse, e error) {

	if count < 10 {
		count = 10
	}

	lst = make([]*warehouse.Warehouse, 0, count)

	if page < 0 {
		page = 0
	}

	offset := page * count

	rows, e := r.queryWarehousesStmt.Query(count, offset)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		w := &warehouse.Warehouse{}
		if e = rows.Scan(&w.ID, &w.Address, &w.Capacity); e != nil {
			log.Println("GetWarehouses", "Error get Warehouses [", e, "]")
			return nil, fmt.Errorf("error get Warehouses %w", e)
		}
		lst = append(lst, w)
	}

	return lst, nil
}

func (r *warehouseRepo) GetWarehousesCnt(ctx context.Context) (total int, e error) {
	if e = r.getWarehousesCntStmt.QueryRow().Scan(&total); e != nil {
		log.Println("GetWarehousesCnt", "Error get Warehouses total [", e, "]")
		return 0, fmt.Errorf("error get Warehouses total %w", e)
	}
	return total, e
}

func (r *warehouseRepo) GetWarehouse(ctx context.Context, id int64) (*warehouse.Warehouse, error) {
	var w warehouse.Warehouse
	e := r.getWarehouseStmt.QueryRow(id).Scan(&w.ID, &w.Address, &w.Capacity)
	if e != nil {
		log.Println("GetWarehouse", "Error get Warehouse [", e, "]")
		return nil, fmt.Errorf("error get Warehouse %w", e)
	}

	return &w, nil
}

func (r *warehouseRepo) GetWarehouseBooks(ctx context.Context, id, page, count int) (lst []*warehouse.WarehouseBooks, e error) {

	if count < 10 {
		count = 10
	}

	lst = make([]*warehouse.WarehouseBooks, 0, count)

	if page < 0 {
		page = 0
	}

	offset := page * count

	rows, e := r.queryWarehouseBooksStmt.Query(id, count, offset)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		w := &warehouse.WarehouseBooks{}
		if e = rows.Scan(&w.ISBN, &w.Title, &w.Author, &w.Genre, &w.Price, &w.Publisher, &w.QuantityOnStock); e != nil {
			log.Println("GetWarehouseBooks", "Error get WarehouseBooks [", e, "]")
			return nil, fmt.Errorf("error get WarehouseBooks %w", e)
		}
		lst = append(lst, w)
	}

	return lst, nil
}

func (r *warehouseRepo) GetWarehouseBooksCnt(ctx context.Context, id int) (total int, e error) {
	if e = r.getWarehouseBooksCntStmt.QueryRow(id).Scan(&total); e != nil {
		log.Println("GetWarehousesCnt", "Error get Warehouses total [", e, "]")
		return 0, fmt.Errorf("error get Warehouses total %w", e)
	}
	return total, e
}