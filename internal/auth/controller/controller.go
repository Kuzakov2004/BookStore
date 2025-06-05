package controller

import (
	"BookStore/internal/auth/service"
	controller2 "BookStore/pkg/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type controller struct {
	srvc service.AuthAdminService
}

func NewBookController(s service.AuthAdminService) (controller2.HttpController, error) {
	return &controller{
		srvc: s,
	}, nil
}

func (c *controller) Init(r *gin.RouterGroup) error {
	bg := r.Group("/admin")
	bg.GET("/login", c.login)
	bg.POST("/login", c.postLogin)
	bg.GET("/logout", c.logout)

	return nil
}

func (c *controller) login(gc *gin.Context) {
	gc.HTML(200, "admin.tpl", gin.H{
		"title": "Admin login",
	})
}
func (c *controller) postLogin(gc *gin.Context) {
	session := sessions.Default(c)

	login := gc.PostForm("login")
	pass := gc.PostForm("pass")

	// Validate form input
	if strings.Trim(login, " ") == "" || strings.Trim(pass, " ") == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	if c.srvc.Lo {
		gc.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set("auid", auid) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	gc.Redirect(http.StatusFound, "/")
}

func (c *controller) logout(gc *gin.Context) {

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
