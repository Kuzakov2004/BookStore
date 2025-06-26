package controller

import (
	"BookStore/internal/auth"
	"BookStore/internal/order"
	"BookStore/internal/order/service"
	controller2 "BookStore/pkg/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type controller struct {
	srvc service.OrderService
}

func NewOrderController(s service.OrderService) (controller2.HttpController, error) {
	return &controller{
		srvc: s,
	}, nil
}

func (c *controller) Init(r *gin.RouterGroup) error {
	og := r.Group("/admin/order")
	og.Use(auth.AdminAuthRequired)
	og.GET("/", c.orders)
	og.GET("/:id", c.order)

	og.GET("/create", c.createOrder)

	og.POST("/selectclient", c.selectClient)
	og.GET("/selectclient", c.selectClient)

	og.POST("/selectclientdone", c.selectClientDone)
	og.POST("/saveship", c.saveShip)

	og.POST("/selectbook", c.selectBook)
	og.GET("/selectbook", c.selectBook)
	og.POST("/selectbookdone", c.selectBookDone)

	og.POST("/saveqty", c.saveQty)

	og.GET("/:id/pay", c.pay)
	og.GET("/:id/send", c.send)

	og.GET("/:id/edit", c.editOrder)
	og.GET("/:id/delbook", c.delBookFromOrder)

	return nil
}

func (c *controller) orders(gc *gin.Context) {

	status := gc.DefaultQuery("status", "N")
	if len(status) == 0 {
		status = "N"
	}

	orders, cnt, e := c.srvc.GetOrders(gc.Request.Context(), status, 0, 50)

	activeActive := ""
	activePayed := ""
	activeCompleted := ""

	if status == "N" {
		activeActive = "active"
	}
	if status == "P" {
		activePayed = "active"
	}
	if status == "S" {
		activeCompleted = "active"
	}

	if e != nil {
		log.Println("Error get orders", e)
	}

	gc.HTML(200, "orders.tpl", gin.H{
		"title":           "Список заказов",
		"orders":          orders,
		"cnt":             cnt,
		"activeActive":    activeActive,
		"activePayed":     activePayed,
		"activeCompleted": activeCompleted,
	})
}

func (c *controller) order(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)

	order, e := c.srvc.GetOrder(gc.Request.Context(), id)
	if e != nil {
		log.Println("Error get order", e)
	}

	gc.HTML(200, "order.tpl", gin.H{
		"title": order.Name,
		"order": order,
	})
}

func (c *controller) createOrder(gc *gin.Context) {
	id, e := c.srvc.CreateOrder(gc.Request.Context())
	if e != nil {
		log.Println("Error create order", e)
	}

	gc.Redirect(http.StatusFound, "/admin/order/"+strconv.FormatInt(id, 10)+"/edit")
}

func (c *controller) editOrder(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)

	if id <= 0 {
		gc.Redirect(http.StatusFound, "/admin/order")
	}

	order, e := c.srvc.GetOrder(gc.Request.Context(), id)
	if e != nil {
		log.Println("Error edit order", e)
	}

	gc.HTML(200, "editorder.tpl", gin.H{
		"title": order.Name,
		"order": order,
	})
}

func (c *controller) selectClient(gc *gin.Context) {

	s := gc.PostForm("str")
	order := gc.PostForm("order")

	str := "%" + s + "%"

	clients, e := c.srvc.FindClient(gc.Request.Context(), str)
	if e != nil {
		log.Println("Error get clients", e)
	}

	gc.HTML(200, "admin/selectclient.tpl", gin.H{
		"title":   "Выбор клиента",
		"clients": clients,
		"str":     s,
		"orderID": order,
	})
}

func (c *controller) selectBook(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.PostForm("order"), 10, 64)

	if orderId == 0 {
		gc.Redirect(http.StatusFound, "/admin/order")
	}

	books, e := c.srvc.FindBook(gc.Request.Context(), orderId, 0, 50)
	if e != nil {
		log.Println("Error get books", e)
	}

	gc.HTML(200, "admin/selectbook.tpl", gin.H{
		"title":   "Выбор клиента",
		"books":   books,
		"orderID": orderId,
	})
}

func (c *controller) selectBookDone(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.PostForm("order"), 10, 64)

	if orderId == 0 {
		gc.Redirect(http.StatusFound, "/admin/order")
	}

	ids := gc.PostFormArray("book_id")
	if len(ids) == 0 {
		gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit#books")
	}

	books := make([]int64, len(ids))
	for i, b := range ids {
		books[i], _ = strconv.ParseInt(b, 10, 64)
	}

	e := c.srvc.AddBooks(gc.Request.Context(), orderId, books)
	if e != nil {
		log.Println("Error add books", e)
	}

	gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit#books")
}

func (c *controller) selectClientDone(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.PostForm("order"), 10, 64)
	clientId, _ := strconv.ParseInt(gc.PostForm("client_id"), 10, 64)

	e := c.srvc.SetOrderClient(gc.Request.Context(), orderId, clientId)
	if e != nil {
		log.Println("Error set client", e)
	}

	gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit#ship")
}

func (c *controller) saveShip(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.PostForm("order"), 10, 64)

	var ship order.Ship

	if e := gc.ShouldBind(&ship); e != nil {
		o, _ := c.srvc.GetOrder(gc.Request.Context(), orderId)
		if e != nil {
			gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit#ship")
		}
		gc.HTML(200, "admin/editorder.tpl", gin.H{
			"title": "Новая книга",
			"order": o,
			"err":   e.Error(),
		})
		return
	}

	e := c.srvc.SaveShip(gc.Request.Context(), orderId, &ship)
	if e != nil {
		log.Println("Error set ship", e)
	}

	gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit#books")
}

func (c *controller) saveQty(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.PostForm("order"), 10, 64)
	pay := gc.PostForm("pay")

	ids := gc.PostFormArray("book_id")
	if len(ids) == 0 {
		gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit#books")
	}

	q := gc.PostFormArray("qty")
	if len(q) == 0 || len(q) != len(ids) {
		gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit#books")
	}

	books := make([]int64, len(ids))
	qty := make([]int, len(q))

	for i, id := range ids {
		books[i], _ = strconv.ParseInt(id, 10, 64)
		qty[i], _ = strconv.Atoi(q[i])
	}

	var e error
	if pay != "pay" {
		e = c.srvc.SaveBookQty(gc.Request.Context(), orderId, books, qty)
		gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit#books")
	} else {
		e = c.srvc.Pay(gc.Request.Context(), orderId, books, qty)
		gc.Redirect(http.StatusFound, "/admin/order?status=P")
	}
	if e != nil {
		log.Println("Error set ship", e)
	}

}

func (c *controller) pay(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.Param("id"), 10, 64)

	e := c.srvc.Pay(gc.Request.Context(), orderId, []int64{}, []int{})
	if e != nil {
		log.Println("Error set ship", e)
	}
	gc.Redirect(http.StatusFound, "/admin/order/"+gc.Param("id"))
}

func (c *controller) send(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.Param("id"), 10, 64)

	e := c.srvc.Send(gc.Request.Context(), orderId)
	if e != nil {
		log.Println("Error set ship", e)
	}
	gc.Redirect(http.StatusFound, "/admin/order?status=S")
}

func (c *controller) delBookFromOrder(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.Param("id"), 10, 64)
	bookId, _ := strconv.ParseInt(gc.Query("book"), 10, 64)

	e := c.srvc.DelBookFromOrder(gc.Request.Context(), orderId, bookId)
	if e != nil {
		log.Println("Error del book from order", e)
	}
	gc.Redirect(http.StatusFound, "/admin/order/"+gc.Param("id")+"/edit#books")
}
