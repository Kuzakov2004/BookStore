package book

type Book struct {
	ID              int64   `json:"id"`
	ISBN            string  `json:"isbn"`
	Title           string  `json:"title"`
	Price           float32 `json:"price"`
	Publisher       string  `json:"publisher"`
	Author          string  `json:"author"`
	PublicationYear int     `json:"publication_year"`
	Genre           string  `json:"genre"`
}

type FullInfo struct {
	Book
	Descr string `json:"descr"`
	Qty   int    `json:"qty"`
}
