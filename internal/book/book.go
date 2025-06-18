package book

type Book struct {
	ID              int64   `form:"id" json:"id" binding:"-"`
	ISBN            string  `form:"isbn" json:"isbn" binding:"required"`
	Title           string  `form:"title" json:"title" binding:"required"`
	Price           float32 `form:"price" json:"price" binding:"required"`
	Publisher       string  `form:"publisher" json:"publisher" binding:"-"`
	Author          string  `form:"author" json:"author" binding:"-"`
	PublicationYear int     `form:"publication_year" json:"publication_year" binding:"required"`
	Genre           string  `form:"genre" json:"genre" binding:"required"`
}

type FullInfo struct {
	Book
	Descr string `form:"descr" json:"descr" binding:"-"`
	Qty   int    `form:"qty" json:"qty" binding:"-"`
}
