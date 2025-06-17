package controller

import (
	"BookStore/internal/auth/service"
	controller2 "BookStore/pkg/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

type controller struct {
	srvc service.AuthService
}

func NewAuthController(s service.AuthService) (controller2.HttpController, error) {
	return &controller{
		srvc: s,
	}, nil
}

func (c *controller) Init(r *gin.RouterGroup) error {
	bg := r.Group("/admin")
	bg.GET("/login", c.login)
	bg.POST("/login", c.postLogin)
	bg.GET("/logout", c.logout)

	cg := r.Group("/client")
	cg.GET("/login", c.clogin)
	cg.POST("/login", c.cpostLogin)
	cg.GET("/logout", c.clogout)

	return nil
}

func (c *controller) login(gc *gin.Context) {
	session := sessions.Default(gc)
	user := session.Get("auid")
	if user != nil {
		gc.Redirect(http.StatusFound, "/admin/books")
		return
	}

	gc.HTML(200, "admin.tpl", gin.H{
		"title": "Admin login",
	})
}
func (c *controller) clogin(gc *gin.Context) {
	session := sessions.Default(gc)
	user := session.Get("cuid")
	if user != nil {
		gc.Redirect(http.StatusFound, "/")
		return
	}

	gc.HTML(200, "client.tpl", gin.H{
		"title": "Client login",
	})
}
func (c *controller) postLogin(gc *gin.Context) {
	session := sessions.Default(gc)

	login := gc.PostForm("login")
	pass := gc.PostForm("pass")

	// Validate form input
	if strings.Trim(login, " ") == "" || strings.Trim(pass, " ") == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	auid, e := c.srvc.Login(gc.Request.Context(), login, pass)
	if e != nil {
		log.Println("Error login ", e)
		gc.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set("auid", auid) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		log.Println("Error login ", e)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	gc.Redirect(http.StatusFound, "/admin/books")
}

func (c *controller) cpostLogin(gc *gin.Context) {
	session := sessions.Default(gc)

	login := gc.PostForm("login")
	pass := gc.PostForm("pass")

	// Validate form input
	if strings.Trim(login, " ") == "" || strings.Trim(pass, " ") == "" {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	// Check for username and password match, usually from a database
	cuid, e := c.srvc.Login(gc.Request.Context(), login, pass)
	if e != nil {
		log.Println("Error login ", e)
		gc.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	// Save the username in the session
	session.Set("cuid", cuid) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		log.Println("Error login ", e)
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	gc.Redirect(http.StatusFound, "/")
}

func (c *controller) logout(gc *gin.Context) {

	session := sessions.Default(gc)
	user := session.Get("auid")
	if user == nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}

	session.Delete("auid")
	if err := session.Save(); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	gc.Redirect(http.StatusFound, "/")
}

func (c *controller) clogout(gc *gin.Context) {
	session := sessions.Default(gc)
	user := session.Get("cuid")
	if user == nil {
		gc.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete("cuid")
	if err := session.Save(); err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	gc.Redirect(http.StatusFound, "/")
}
