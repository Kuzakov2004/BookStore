package order

type Ship struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
	City    string `form:"city" json:"city" binding:"required"`
	ZipCode string `form:"zip" json:"zip_code" binding:"required"`
	Country string `form:"country" json:"country" `
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
	Status string `json:"status"`
}

type OrderItem struct {
	OrderID    int64   `json:"order_id"`
	BookID     int64   `json:"book_id"`
	BookTitle  string  `json:"book_title"`
	BookAuthor string  `json:"book_author"`
	Price      float64 `json:"price"`
	Qty        int     `json:"qty"`
	InStock    int     `json:"in_stock"`
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
