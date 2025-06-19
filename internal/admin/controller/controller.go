package controller

import (
	"BookStore/internal/admin/service"
	"BookStore/internal/auth"
	"BookStore/internal/book"
	service2 "BookStore/internal/book/service"
	"BookStore/internal/config"
	"BookStore/internal/publisher"
	service3 "BookStore/internal/publisher/service"
	controller2 "BookStore/pkg/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path"
	"strconv"
)

type controller struct {
	cfg           *config.Config
	adminSrvc     service.AdminService
	bookSrvc      service2.BookService
	publisherSrvc service3.PublisherService
}

func NewAdminController(cfg *config.Config, s service.AdminService, bs service2.BookService, ps service3.PublisherService) (controller2.HttpController, error) {
	return &controller{
		cfg:           cfg,
		adminSrvc:     s,
		bookSrvc:      bs,
		publisherSrvc: ps,
	}, nil
}

func (c *controller) Init(r *gin.RouterGroup) error {
	bg := r.Group("/admin")
	bg.Use(auth.AdminAuthRequired)
	bg.GET("/book", c.books)

	bg.GET("/book/:id/edit", c.bookEdit)
	bg.POST("/book/:id/edit", c.postBookEdit)

	bg.GET("/book/create", c.bookCreate)
	bg.POST("/book/create", c.postBookCreate)

	bg.GET("/book/:id/delete", c.bookDelete)

	bg.GET("/publisher", c.publishers)

	bg.GET("/publisher/:id/edit", c.publisherEdit)
	bg.POST("/publisher/:id/edit", c.postPublisherEdit)

	bg.GET("/publisher/create", c.publisherCreate)
	bg.POST("/publisher/create", c.postPublisherCreate)

	bg.GET("/publisher/:id/delete", c.publisherDelete)

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
		"title":    book.Title,
		"book":     book,
		"isCreate": false,
	})
}

func (c *controller) bookCreate(gc *gin.Context) {

	gc.HTML(200, "admin/bookedit.tpl", gin.H{
		"title":    "Новая книга",
		"book":     book.FullInfo{},
		"isCreate": true,
	})
}

func (c *controller) postBookEdit(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)
	var bookInfo book.FullInfo
	bookInfo.ID = id

	if e := gc.ShouldBind(&bookInfo); e != nil {
		gc.HTML(200, "admin/bookedit.tpl", gin.H{
			"title": "Новая книга",
			"book":  bookInfo,
			"err":   e.Error(),
		})
		return
	}

	e := c.adminSrvc.UpdateBook(gc.Request.Context(), &bookInfo)

	if e != nil {
		log.Println("Error update book", e)
		gc.HTML(200, "admin/bookedit.tpl", gin.H{
			"title": "Новая книга",
			"book":  bookInfo,
			"err":   e.Error(),
		})
		return
	}

	file, e := gc.FormFile("image")
	if file != nil && (file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png") {
		e = fmt.Errorf("Invalid image mime type %s", file.Header.Get("Content-Type"))
		gc.HTML(200, "admin/bookedit.tpl", gin.H{
			"title": "Новая книга",
			"book":  bookInfo,
			"err":   e.Error(),
		})
		return
	}

	if file != nil {
		filename := path.Join(c.cfg.ImagePath, strconv.FormatInt(id, 10)+".jpg")
		if err := gc.SaveUploadedFile(file, filename); err != nil {
			log.Println("upload file err: ", err.Error())
		}
	}

	gc.Redirect(http.StatusFound, "/admin/book")
}

func (c *controller) postBookCreate(gc *gin.Context) {

	var bookInfo book.FullInfo

	if e := gc.ShouldBind(&bookInfo); e != nil {
		gc.HTML(200, "admin/bookedit.tpl", gin.H{
			"title": "Новая книга",
			"book":  bookInfo,
			"err":   e.Error(),
		})
		return
	}

	id, e := c.adminSrvc.CreateBook(gc.Request.Context(), &bookInfo)

	if e != nil {
		log.Println("Error create book", e)
		gc.HTML(200, "admin/bookedit.tpl", gin.H{
			"title": "Новая книга",
			"book":  bookInfo,
			"err":   e.Error(),
		})
		return
	}

	file, e := gc.FormFile("image")

	if file != nil && (file.Header.Get("Content-Type") != "image/jpeg" && file.Header.Get("Content-Type") != "image/png") {
		e = fmt.Errorf("Invalid image mime type %s", file.Header.Get("Content-Type"))
		gc.HTML(200, "admin/bookedit.tpl", gin.H{
			"title": "Новая книга",
			"book":  bookInfo,
			"err":   e.Error(),
		})
		return
	}

	if file != nil {
		filename := path.Join(c.cfg.ImagePath, strconv.FormatInt(id, 10)+".jpg")
		if err := gc.SaveUploadedFile(file, filename); err != nil {
			log.Println("upload file err: ", err.Error())
		}
	}

	gc.Redirect(http.StatusFound, "/admin/book")
}

func (c *controller) bookDelete(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)

	e := c.adminSrvc.DeleteBook(gc.Request.Context(), id)

	if e != nil {
		log.Println("Error delete book", e)
	}

	gc.Redirect(http.StatusFound, "/admin/book")
}

func (c *controller) publishers(gc *gin.Context) {

	publishers, cnt, e := c.publisherSrvc.GetPublishers(gc.Request.Context(), 0, 50)

	if e != nil {
		log.Println("Error get publishers", e)
	}

	gc.HTML(200, "admin/publishers.tpl", gin.H{
		"title":      "Список издательств",
		"publishers": publishers,
		"cnt":        cnt,
		"isAdmin":    gc.Keys["isAdmin"],
	})
}

func (c *controller) publisherEdit(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)
	publisher, e := c.publisherSrvc.GetPublisher(gc.Request.Context(), id)

	if e != nil {
		log.Println("Error get publisher", e)
	}

	gc.HTML(200, "admin/publisheredit.tpl", gin.H{
		"title":     publisher.Name,
		"publisher": publisher,
		"isCreate":  false,
	})
}

func (c *controller) publisherCreate(gc *gin.Context) {

	gc.HTML(200, "admin/publisheredit.tpl", gin.H{
		"title":     "Новый издатель",
		"publisher": publisher.Publisher{},
		"isCreate":  true,
	})
}

func (c *controller) postPublisherEdit(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)
	var p publisher.Publisher
	p.ID = id

	if e := gc.ShouldBind(&p); e != nil {
		gc.HTML(200, "admin/publisheredit.tpl", gin.H{
			"title":     "Новый издатель",
			"publisher": p,
			"err":       e.Error(),
		})
		return
	}

	e := c.adminSrvc.UpdatePublisher(gc.Request.Context(), &p)

	if e != nil {
		log.Println("Error update publisher", e)
		gc.HTML(200, "admin/publisheredit.tpl", gin.H{
			"title":     "Новый издатель",
			"publisher": p,
			"err":       e.Error(),
		})
		return
	}

	gc.Redirect(http.StatusFound, "/admin/publisher")
}

func (c *controller) postPublisherCreate(gc *gin.Context) {

	var p publisher.Publisher

	if e := gc.ShouldBind(&p); e != nil {
		gc.HTML(200, "admin/publisheredit.tpl", gin.H{
			"title":     "Новый издатель",
			"publisher": p,
			"err":       e.Error(),
		})
		return
	}

	_, e := c.adminSrvc.CreatePublisher(gc.Request.Context(), &p)

	if e != nil {
		log.Println("Error create book", e)
		gc.HTML(200, "admin/publisheredit.tpl", gin.H{
			"title":     "Новая книга",
			"publisher": p,
			"err":       e.Error(),
		})
		return
	}

	gc.Redirect(http.StatusFound, "/admin/publisher")
}

func (c *controller) publisherDelete(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)

	e := c.adminSrvc.DeletePublisher(gc.Request.Context(), id)

	if e != nil {
		log.Println("Error delete publisher", e)
	}

	gc.Redirect(http.StatusFound, "/admin/publisher")
}
