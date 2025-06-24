package repo

import (
	"BookStore/internal/order"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type orderRepo struct {
	connPool *sql.DB

	queryOrdersStmt     *sql.Stmt
	getOrdersCntStmt    *sql.Stmt
	getOrderStmt        *sql.Stmt
	queryOrderItemsStmt *sql.Stmt
	createOrderStmt     *sql.Stmt

	findClientStmt *sql.Stmt
}

func NewOrderRepo(db *sql.DB) (OrderRepo, error) {

	r := &orderRepo{connPool: db}
	var e error

	r.queryOrdersStmt, e = db.Prepare(`SELECT o.id, COALESCE(o.client_id, 0), COALESCE(c.first_name || ' ' || c.last_name,'') , COALESCE(c.phone, ''), COALESCE(o.amount, 0), 
       										COALESCE(o.dt, now()), 
											COALESCE((select sum(od.qty) from store.order_items od where od.order_id = o.id), 0) qty, COALESCE(o.ship_address, '')
											FROM store.orders o 
											LEFT JOIN store.clients c ON (o.client_id = c.id)
											WHERE o.status = $1
    										LIMIT $2 OFFSET $3`)
	if e != nil {
		return nil, e
	}

	r.getOrdersCntStmt, e = db.Prepare(`SELECT count(*) FROM store.orders o WHERE o.status = $1`)
	if e != nil {
		return nil, e
	}

	r.queryOrderItemsStmt, e = db.Prepare(`SELECT oi.order_id, oi.book_id, 
													b.title,
														case when a.middle_name is null then a.first_name || ' ' || a.last_name
															 else substr(a.first_name, 1, 1) || '. ' || substr(a.middle_name, 1, 1) || '. ' || a.last_name
														end author,     
													oi.item_price, oi.qty
												FROM store.order_items oi 
												JOIN store.books b ON (b.id = oi.book_id)
												JOIN store.authors a ON (b.author_id = a.id)
													WHERE oi.order_id = $1 ORDER BY b.title`)
	if e != nil {
		return nil, e
	}

	r.getOrderStmt, e = db.Prepare(`SELECT o.id, coalesce(o.client_id, 0), COALESCE(c.first_name || ' ' || c.last_name, ''), COALESCE(c.phone,''), COALESCE(o.amount, 0), 
       										COALESCE(o.dt, now()),
											COALESCE(o.ship_name, ''), COALESCE(o.ship_address, ''), COALESCE(o.ship_city, ''), COALESCE(o.ship_zip_code, ''),
       										COALESCE(o.ship_country, '')
											FROM store.orders o 
											LEFT JOIN store.clients c ON (o.client_id = c.id)
											WHERE o.id = $1`)
	if e != nil {
		return nil, e
	}

	r.findClientStmt, e = db.Prepare(`SELECT c.id, c.first_name, c.last_name, c.middle_name, c.login, c.phone
											FROM store.clients c 
											WHERE c.first_name like $1 or c.last_name like $1 or c.middle_name like $1  or c.middle_name like $1  or c.login like $1
											LIMIT 10`)
	if e != nil {
		return nil, e
	}

	r.createOrderStmt, e = db.Prepare(`INSERT INTO store.orders (client_id) values (null) RETURNING id`)
	if e != nil {
		return nil, e
	}

	return r, nil
}

func (r *orderRepo) FindClient(ctx context.Context, str string) (lst []*order.Client, e error) {
	lst = make([]*order.Client, 0, 10)

	rows, e := r.findClientStmt.Query(str)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		//c.id, c.first_name, c.last_name, c.middle_name, c.login, c.phone
		c := &order.Client{}
		if e = rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.MiddleName, &c.Login, &c.Phone); e != nil {
			log.Println("FindClient", "Error find client [", e, "]")
			return nil, fmt.Errorf("error find client %w", e)
		}
		lst = append(lst, c)
	}
	return lst, nil
}

func (r *orderRepo) GetOrders(ctx context.Context, status string, page, count int) (lst []*order.Order, e error) {
	lst = make([]*order.Order, 0, count)

	if count < 10 {
		count = 10
	}

	if page < 0 {
		page = 0
	}

	offset := page * count

	rows, e := r.queryOrdersStmt.Query(status, count, offset)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		//SELECT o.id, o.client_id, fio, c.phone, o.amount, o.dt,
		//qty, o.ship_address
		o := &order.Order{}
		if e = rows.Scan(&o.ID, &o.ClientID, &o.ClientFIO, &o.ClientPhone, &o.Amount, &o.Dt, &o.Qty, &o.Ship.Address); e != nil {
			log.Println("GetOrders", "Error get orders [", e, "]")
			return nil, fmt.Errorf("error get orders %w", e)
		}
		lst = append(lst, o)
	}
	return lst, nil
}

func (r *orderRepo) GetOrdersCnt(ctx context.Context, status string) (total int, e error) {
	if e = r.getOrdersCntStmt.QueryRow(status).Scan(&total); e != nil {
		log.Println("GetOrdersCnt", "Error get orders total [", e, "]")
		return 0, fmt.Errorf("error get orders total %w", e)
	}
	return total, e
}

func (r *orderRepo) GetOrder(ctx context.Context, id int64) (*order.OrderDetail, error) {
	var o order.OrderDetail

	//o.id, o.client_id, c.first_name || ' ' || c.last_name, c.phone, o.amount, o.dt
	//o.ship_name, o.ship_address, o.ship_city, o.ship_zip_code, o.ship_country

	e := r.getOrderStmt.QueryRow(id).Scan(&o.ID, &o.ClientID, &o.ClientFIO, &o.ClientPhone, &o.Amount, &o.Dt,
		&o.Ship.Name, &o.Ship.Address, &o.Ship.City, &o.Ship.ZipCode, &o.Ship.Country)
	if e != nil {
		log.Println("GetOrder", "Error get order [", e, "]")
		return nil, fmt.Errorf("error get order %w", e)
	}

	o.Items = make([]*order.OrderItem, 0, 10)

	rows, e := r.queryOrderItemsStmt.Query(id)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		//oi.order_id, oi.book_id, b.title, author, oi.item_price, oi.qty
		oi := &order.OrderItem{}
		if e = rows.Scan(&oi.OrderID, &oi.BookID, &oi.BookTitle, &oi.BookAuthor, &oi.Price, &oi.Qty); e != nil {
			log.Println("GetOrder", "Error get order items [", e, "]")
			return nil, fmt.Errorf("error get order items %w", e)
		}
		o.Items = append(o.Items, oi)
	}

	return &o, nil
}

func (r *orderRepo) CreateOrder(ctx context.Context) (int64, error) {
	var id int64
	e := r.createOrderStmt.QueryRow().Scan(&id)
	if e != nil {
		log.Println("CreateOrder", "Error create order [", e, "]")
		return 0, fmt.Errorf("error create order %w", e)
	}

	return id, nil
}
