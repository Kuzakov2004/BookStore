package controller

import (
	"BookStore/internal/auth"
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

	og.GET("/:id/edit", c.editOrder)

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
	if status == "C" {
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
func (c *controller) selectClientDone(gc *gin.Context) {

	orderId, _ := strconv.ParseInt(gc.PostForm("order"), 10, 64)
	clientId, _ := strconv.ParseInt(gc.PostForm("client_id"), 10, 64)

	e := c.srvc.SetOrderClient(gc.Request.Context(), orderId, clientId)
	if e != nil {
		log.Println("Error set client", e)
	}

	gc.Redirect(http.StatusFound, "/admin/order/"+gc.PostForm("order")+"/edit")
}
