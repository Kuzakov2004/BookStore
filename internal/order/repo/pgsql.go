package repo

import (
	"BookStore/internal/book"
	"BookStore/internal/order"
	"context"
	"database/sql"
	"fmt"
	"log"
)

type wrhsBookQty struct {
	wrhsId int64
	bookId int64
	qty    int
}

type orderRepo struct {
	connPool *sql.DB

	queryOrdersStmt       *sql.Stmt
	getOrdersCntStmt      *sql.Stmt
	getOrderStmt          *sql.Stmt
	queryOrderItemsStmt   *sql.Stmt
	createOrderStmt       *sql.Stmt
	setOrderClientStmt    *sql.Stmt
	saveShipStmt          *sql.Stmt
	addBookStmt           *sql.Stmt
	saveBookQtyStmt       *sql.Stmt
	updateStatusStmt      *sql.Stmt
	delBookFromOrderStmt  *sql.Stmt
	updateOrderAmountStmt *sql.Stmt

	updateWrhsQtyStmt  *sql.Stmt
	queryWrhsBooksStmt *sql.Stmt

	findClientStmt *sql.Stmt
	findBookStmt   *sql.Stmt
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
													oi.item_price, oi.qty,
       												COALESCE((SELECT SUM(wb.qty) FROM store.warehouse_books wb WHERE wb.book_id = oi.book_id), 0)
												FROM store.order_items oi 
												JOIN store.books b ON (b.id = oi.book_id)
												JOIN store.authors a ON (b.author_id = a.id)
													WHERE oi.order_id = $1 ORDER BY b.title`)
	if e != nil {
		return nil, e
	}

	r.getOrderStmt, e = db.Prepare(`SELECT o.id, coalesce(o.client_id, 0), COALESCE(c.first_name || ' ' || c.last_name, ''), COALESCE(c.phone,''), COALESCE(o.amount, 0), 
       										COALESCE(o.dt, now()), COALESCE(o.qty),
											COALESCE(o.ship_name, ''), COALESCE(o.ship_address, ''), COALESCE(o.ship_city, ''), COALESCE(o.ship_zip_code, ''),
       										COALESCE(o.ship_country, ''), o.status
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

	r.setOrderClientStmt, e = db.Prepare(`update store.orders set client_id= $2 where id = $1`)
	if e != nil {
		return nil, e
	}

	r.updateOrderAmountStmt, e = db.Prepare(`update store.orders set 
                        amount=(select sum(oi.item_price*oi.qty) from store.order_items oi where oi.order_id=$1),
                        qty=(select sum(oi.qty) from store.order_items oi where oi.order_id=$1) where id = $1`)
	if e != nil {
		return nil, e
	}

	r.saveShipStmt, e = db.Prepare(`update store.orders set ship_name=$2, ship_address=$3, ship_city=$4, ship_zip_code=$5, ship_country=$6 where id = $1`)
	if e != nil {
		return nil, e
	}

	r.delBookFromOrderStmt, e = db.Prepare(`delete from store.order_items where order_id = $1 and book_id = $2`)
	if e != nil {
		return nil, e
	}

	r.findBookStmt, e = db.Prepare(`SELECT b.id, b.isbn, b.title, b.price, p.name publisher, a.first_name || ' ' || a.last_name author, b.publication_year, b.genre 
											FROM store.books b 
											JOIN store.publishers p ON (b.publisher_id = p.id)
											JOIN store.authors a ON (b.author_id = a.id)
											WHERE b.id not in (SELECT oi.book_id FROM store.order_items oi WHERE oi.order_id=$1) ORDER BY b.title DESC
    										LIMIT $2 OFFSET $3`)

	r.addBookStmt, e = db.Prepare(`INSERT INTO store.order_items (order_id, book_id, item_price, qty) 
										 VALUES ($1, $2, (SELECT b.price FROM store.books b where b.id=$2), 1)`)
	if e != nil {
		return nil, e
	}

	r.saveBookQtyStmt, e = db.Prepare(`UPDATE store.order_items SET qty = $3 WHERE order_id=$1 and book_id=$2`)
	if e != nil {
		return nil, e
	}

	r.updateStatusStmt, e = db.Prepare(`UPDATE store.orders SET status = $2 WHERE id=$1`)
	if e != nil {
		return nil, e
	}

	r.updateWrhsQtyStmt, e = db.Prepare(`UPDATE store.warehouse_books SET qty = qty - $3 WHERE wrhs_id=$1 and book_id=$2`)
	if e != nil {
		return nil, e
	}

	r.queryWrhsBooksStmt, e = db.Prepare(`SELECT wb.wrhs_id, wb.book_id, wb.qty
											FROM store.warehouse_books wb 
											WHERE wb.book_id=$1`)
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

func (r *orderRepo) getWrhsQty(ctx context.Context, id int64) (lst []*wrhsBookQty, e error) {
	lst = make([]*wrhsBookQty, 0, 50)

	rows, e := r.queryWrhsBooksStmt.Query(id)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		w := &wrhsBookQty{}
		if e = rows.Scan(&w.wrhsId, &w.bookId, &w.qty); e != nil {
			log.Println("getWrhsQty", "Error get getWrhsQty [", e, "]")
			return nil, fmt.Errorf("error get getWrhsQty %w", e)
		}
		lst = append(lst, w)
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

	e := r.getOrderStmt.QueryRow(id).Scan(&o.ID, &o.ClientID, &o.ClientFIO, &o.ClientPhone, &o.Amount, &o.Dt, &o.Qty,
		&o.Ship.Name, &o.Ship.Address, &o.Ship.City, &o.Ship.ZipCode, &o.Ship.Country, &o.Status)
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
		//oi.order_id, oi.book_id, b.title, author, oi.item_price, oi.qty, oi.instock
		oi := &order.OrderItem{}
		if e = rows.Scan(&oi.OrderID, &oi.BookID, &oi.BookTitle, &oi.BookAuthor, &oi.Price, &oi.Qty, &oi.InStock); e != nil {
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

func (r *orderRepo) SetOrderClient(ctx context.Context, orderId int64, clientId int64) error {
	_, e := r.setOrderClientStmt.Exec(orderId, clientId)
	if e != nil {
		log.Println("SetOrderClient", "Error set client order [", e, "]")
		return fmt.Errorf("error set client order %w", e)
	}

	return nil
}

func (r *orderRepo) DelBookFromOrder(ctx context.Context, orderId int64, bookId int64) error {
	_, e := r.delBookFromOrderStmt.Exec(orderId, bookId)
	if e != nil {
		log.Println("DelBookFromOrder", "Error del book from order [", e, "]")
		return fmt.Errorf("error del book from order %w", e)
	}

	return nil
}

func (r *orderRepo) SaveShip(ctx context.Context, orderId int64, ship *order.Ship) error {
	//ship_name=$2, ship_address=$3, ship_city=$4, ship_zip_code=$5, ship_country=$6 where id = $1
	_, e := r.saveShipStmt.Exec(orderId, ship.Name, ship.Address, ship.City, ship.ZipCode, ship.Country)
	if e != nil {
		log.Println("SaveShip", "Error set ship order [", e, "]")
		return fmt.Errorf("error set ship order %w", e)
	}

	return nil
}

func (r *orderRepo) AddBooks(ctx context.Context, orderId int64, books []int64) error {

	for _, id := range books {
		_, e := r.addBookStmt.Exec(orderId, id)
		if e != nil {
			log.Println("AddBooks", "Error add book order [", e, "]")
			return fmt.Errorf("error add book order %w", e)
		}
	}

	_, e := r.updateOrderAmountStmt.Exec(orderId)
	if e != nil {
		log.Println("SaveBookQty", "Error update order amount [", e, "]")
		return fmt.Errorf("error update order amoun %w", e)
	}

	return nil
}

func (r *orderRepo) SaveBookQty(ctx context.Context, orderId int64, books []int64, qty []int) error {

	for i, id := range books {
		_, e := r.saveBookQtyStmt.Exec(orderId, id, qty[i])
		if e != nil {
			log.Println("SaveBookQty", "Error save book qty [", e, "]")
			return fmt.Errorf("error save book qty %w", e)
		}
	}

	_, e := r.updateOrderAmountStmt.Exec(orderId)
	if e != nil {
		log.Println("SaveBookQty", "Error update order amount [", e, "]")
		return fmt.Errorf("error update order amoun %w", e)
	}

	return nil
}

func (r *orderRepo) Pay(ctx context.Context, orderId int64, books []int64, qty []int) error {

	for i, id := range books {
		_, e := r.saveBookQtyStmt.Exec(orderId, id, qty[i])
		if e != nil {
			log.Println("Pay", "Error pay [", e, "]")
			return fmt.Errorf("error pay %w", e)
		}
	}

	_, e := r.updateStatusStmt.Exec(orderId, "P")
	if e != nil {
		log.Println("Pay1", "Error pay [", e, "]")
		return fmt.Errorf("error pay %w", e)
	}

	return nil
}

func (r *orderRepo) Send(ctx context.Context, orderId int64) error {

	o, e := r.GetOrder(ctx, orderId)
	if e != nil {
		log.Println("Send", "Error Send [", e, "]")
		return fmt.Errorf("error Send %w", e)
	}

	for _, it := range o.Items {
		wb, e := r.getWrhsQty(ctx, it.BookID)
		if e != nil {
			log.Println("Send", "Error Send [", e, "]")
			return fmt.Errorf("error Send %w", e)
		}
		for _, w := range wb {
			if w.qty > it.Qty {
				r.updateWrhsQtyStmt.Exec(w.wrhsId, w.bookId, it.Qty)
				break
			} else {
				r.updateWrhsQtyStmt.Exec(w.wrhsId, w.bookId, w.qty)
				it.Qty -= w.qty
			}
		}
	}

	_, e = r.updateStatusStmt.Exec(orderId, "S")
	if e != nil {
		log.Println("Send", "Error send [", e, "]")
		return fmt.Errorf("error send %w", e)
	}

	return nil
}

func (r *orderRepo) FindBook(ctx context.Context, orderId int64, page, count int) (lst []*book.Book, e error) {
	lst = make([]*book.Book, 0, count)

	if count < 10 {
		count = 10
	}

	if page < 0 {
		page = 0
	}
	offset := page * count

	rows, e := r.findBookStmt.Query(orderId, count, offset)
	if e != nil {
		return nil, e
	}

	defer rows.Close()

	for rows.Next() {
		b := &book.Book{}
		if e = rows.Scan(&b.ID, &b.ISBN, &b.Title, &b.Price, &b.Publisher, &b.Author, &b.PublicationYear, &b.Genre); e != nil {
			log.Println("FindBook", "Error get books [", e, "]")
			return nil, fmt.Errorf("error get books %w", e)
		}
		lst = append(lst, b)
	}
	return lst, nil
}
