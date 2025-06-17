package controller

import (
	"BookStore/internal/auth"
	service2 "BookStore/internal/book/service"
	controller2 "BookStore/pkg/controller"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type controller struct {
	bookSrvc service2.BookService
}

func NewAdminController(s service2.BookService) (controller2.HttpController, error) {
	return &controller{
		bookSrvc: s,
	}, nil
}

func (c *controller) Init(r *gin.RouterGroup) error {
	bg := r.Group("/admin")
	bg.Use(auth.AdminAuthRequired)
	bg.GET("/books", c.books)
	bg.GET("/books/:id/edit", c.bookEdit)
	bg.POST("/books/:id/edit", c.postBookEdit)

	return nil
}

func (c *controller) books(gc *gin.Context) {

	books, cnt, e := c.bookSrvc.GetBooks(gc.Request.Context(), "", 0, 50)

	if e != nil {
		log.Println("Error get books", e)
	}

	gc.HTML(200, "admin/books.tpl", gin.H{
		"title":   "Список книг",
		"books":   books,
		"cnt":     cnt,
		"isAdmin": gc.Keys["isAdmin"],
	})
}

func (c *controller) bookEdit(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)
	book, e := c.bookSrvc.GetBook(gc.Request.Context(), id)

	if e != nil {
		log.Println("Error get book", e)
	}

	gc.HTML(200, "admin/bookedit.tpl", gin.H{
		"title": book.Title,
		"book":  book,
	})
}

func (c *controller) postBookEdit(gc *gin.Context) {

	books, cnt, e := c.bookSrvc.GetBooks(gc.Request.Context(), "", 0, 50)

	if e != nil {
		log.Println("Error get books", e)
	}

	gc.HTML(200, "admin/books.tpl", gin.H{
		"title":   "Список книг",
		"books":   books,
		"cnt":     cnt,
		"isAdmin": gc.Keys["isAdmin"],
	})
}
