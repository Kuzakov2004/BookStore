package controller

import (
	"BookStore/internal/book/service"
	controller2 "BookStore/pkg/controller"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type controller struct {
	srvc service.BookService
}

func NewBookController(s service.BookService) (controller2.HttpController, error) {
	return &controller{
		srvc: s,
	}, nil
}

func (c *controller) Init(r *gin.RouterGroup) error {
	bg := r.Group("/book")
	bg.GET("/", c.books)
	bg.GET("/:id", c.book)

	return nil
}

func (c *controller) books(gc *gin.Context) {

	books, cnt, e := c.srvc.GetBooks(gc.Request.Context(), "", 0, 50)

	if e != nil {
		log.Println("Error get books", e)
	}

	gc.HTML(200, "books.tpl", gin.H{
		"title": "Список книг",
		"books": books,
		"cnt":   cnt,
	})
}

func (c *controller) book(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)
	book, e := c.srvc.GetBook(gc.Request.Context(), id)

	if e != nil {
		log.Println("Error get book", e)
	}

	gc.HTML(200, "book.tpl", gin.H{
		"title": book.Title,
		"book":  book,
	})
}
