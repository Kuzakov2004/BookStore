package warehouse

type Warehouse struct {
	ID                 int64  `form:"id" json:"id" binding:"-"`
	Address            string `form:"address" json:"address" binding:"required"`
	Capacity           string `form:"capacity" json:"capacity" binding:"required"`
}

type WarehouseBooks struct {
	//ID            int64  `form:"id" json:"id" binding:"-"`
	ISBN            string  `form:"isbn" json:"isbn" binding:"required"`
	Title           string  `form:"title" json:"title" binding:"required"`
	Author          string  `form:"author" json:"author" binding:"-"`
	Genre           string  `form:"genre" json:"genre" binding:"required"`
	Price           float32 `form:"price" json:"price" binding:"required"`
	Publisher       string  `form:"publisher" json:"publisher" binding:"-"`
	QuantityOnStock int     `form:"quantity_on_stock" json:"quantity_on_stock" binding:"required"`
}