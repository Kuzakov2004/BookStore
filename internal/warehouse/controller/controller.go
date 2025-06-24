package controller

import (
	"log"
	"strconv"

	"BookStore/internal/auth"

	"github.com/gin-gonic/gin"

	service3 "BookStore/internal/warehouse/service"
	controller3 "BookStore/pkg/controller"
)

type controller struct {
	srvc service3.WarehouseService
}

func NewWarehouseController(s service3.WarehouseService) (controller3.HttpController, error) {
	return &controller{
		srvc: s,
	}, nil
}

func (c *controller) Init(r *gin.RouterGroup) error {

	wg := r.Group("/admin/warehouse")
	wg.Use(auth.AdminAuthRequired)
	wg.GET("/", c.warehouses)
	wg.GET("/:id", c.warehouse)
	wg.GET("/:id/books", c.warehouseBooks)

	return nil
}

func (c *controller) warehouses(gc *gin.Context) {

	warehouses, cnt, e := c.srvc.GetWarehouses(gc.Request.Context(), 0, 50)

	if e != nil {
		log.Println("Error get Warehouses", e)
	}

	gc.HTML(200, "warehouses.tpl", gin.H{
		"title":      "Склады",
		"warehouses": warehouses,
		"cnt":        cnt,
		"isAdmin":    true,
	})
}

func (c *controller) warehouseBooks(gc *gin.Context) {
	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)

	// Получаем склад
	warehouse, err := c.srvc.GetWarehouse(gc.Request.Context(), id)
	if err != nil {
		log.Println("Error get Warehouse", err)
		gc.AbortWithStatus(500)
		return
	}

	// Получаем книги на складе
	warehouseBooks, _, e := c.srvc.GetWarehouseBooks(gc.Request.Context(), int(id), 0, 50)
	if e != nil {
		log.Println("Error get Warehouse Books", e)
		gc.AbortWithStatus(500)
		return
	}

	gc.HTML(200, "warehouse_books.tpl", gin.H{
		"title":     "Книги на складе",
		"warehouse": warehouse,
		"books":     warehouseBooks,
		"isAdmin":   true,
	})
}

func (c *controller) warehouse(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)
	warehouse, e := c.srvc.GetWarehouse(gc.Request.Context(), id)

	if e != nil {
		log.Println("Error get warehouse", e)
	}

	gc.HTML(200, "warehouse.tpl", gin.H{
		"title":     warehouse,
		"warehouse": warehouse,
		"isAdmin":   true,
	})
}
