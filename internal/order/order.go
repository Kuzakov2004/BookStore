package order

type Ship struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}
type Order struct {
	ID          int64   `json:"id"`
	ClientID    int64   `json:"client_id"`
	ClientFIO   string  `json:"client_fio"`
	ClientPhone string  `json:"client_phone"`
	Amount      float64 `json:"amount"`
	Dt          string  `json:"dt"`
	Qty         int     `json:"qty"`
	Ship
	Status byte `json:"status"`
}

type OrderItem struct {
	OrderID    int64   `json:"order_id"`
	BookID     int64   `json:"book_id"`
	BookTitle  string  `json:"book_title"`
	BookAuthor string  `json:"book_author"`
	Price      float64 `json:"price"`
	Qty        int     `json:"qty"`
}

type Client struct {
	ID         int64  `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Login      string `json:"login"`
	Phone      string `json:"phone"`
}

type OrderDetail struct {
	Order
	Items []*OrderItem `json:"items"`
}
