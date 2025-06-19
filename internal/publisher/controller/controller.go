package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	service2 "BookStore/internal/publisher/service"
	controller2 "BookStore/pkg/controller"
)

type controller struct {
	srvc service2.PublisherService
}

func NewPublisherController(s service2.PublisherService) (controller2.HttpController, error) {
	return &controller{
		srvc: s,
	}, nil
}

func (c *controller) Init(r *gin.RouterGroup) error {
	bg := r.Group("/publisher")
	bg.GET("/", c.publishers)
	bg.GET("/:id", c.publisher)

	return nil
}

func (c *controller) publishers(gc *gin.Context) {

	publishers, cnt, e := c.srvc.GetPublishers(gc.Request.Context(), 0, 50)

	if e != nil {
		log.Println("Error get publishers", e)
	}

	gc.HTML(200, "publishers.tpl", gin.H{
		"title":      "Издатели",
		"publishers": publishers,
		"cnt":        cnt,
	})
}

func (c *controller) publisher(gc *gin.Context) {

	id, _ := strconv.ParseInt(gc.Param("id"), 10, 64)
	publisher, e := c.srvc.GetPublisher(gc.Request.Context(), id)

	if e != nil {
		log.Println("Error get publisher", e)
	}

	gc.HTML(200, "publisher.tpl", gin.H{
		"title":     publisher.Name,
		"publisher": publisher,
	})
}
