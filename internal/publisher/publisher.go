package publisher

type Publisher struct {
	ID      int64  `form:"id" json:"id" binding:"-"`
	Name    string `form:"name" json:"name" binding:"required"`
	Country string `form:"country" json:"country" binding:"required"`
	Phone   string `form:"phone" json:"phone" binding:"required"`
}
