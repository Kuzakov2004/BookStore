package app

import (
	controller3 "BookStore/internal/admin/controller"
	"BookStore/internal/auth"
	controller2 "BookStore/internal/auth/controller"
	repo2 "BookStore/internal/auth/repo"
	service2 "BookStore/internal/auth/service"
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"BookStore/internal/book/controller"
	"BookStore/internal/book/repo"
	"BookStore/internal/book/service"
	"BookStore/internal/config"
	"BookStore/pkg/app"
)

type storeApp struct {
	cfg         *config.Config
	db          *sql.DB
	srv         *http.Server
	router      *gin.Engine
	bookService service.BookService
}

func NewStoreApp() app.Application {
	app := &storeApp{}

	if e := app.init(); e != nil {
		return nil
	}

	return app
}

func (a *storeApp) Run(ctx context.Context) {
	go func() {
		a.srv.ListenAndServe()
	}()

	<-ctx.Done()

	c, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	a.srv.Shutdown(c)
	<-c.Done()

	time.Sleep(time.Second)
	log.Println("Done")
}

func (a *storeApp) init() error {

	inits := []func() error{
		a.initConfig,
		a.initDb,
		a.initRouter,
		a.initAuth,
		a.initBooks,
		a.initAdmin,
	}

	for _, fn := range inits {
		if e := fn(); e != nil {
			log.Println("Error init app", e.Error())
			return e
		}
	}
	return nil
}

func (a *storeApp) initConfig() error {
	var e error
	a.cfg, e = config.Load()
	return e
}

func (a *storeApp) initDb() error {
	db, e := sql.Open(a.cfg.DB.DriverName, a.cfg.DB.DataSourceName())
	if e != nil {
		fmt.Println("Error init database", e.Error())
		return e
	}

	e = db.Ping()
	if e != nil {
		fmt.Println("Error init database", e.Error())
		return e
	}

	a.db = db
	return nil
}

func (a *storeApp) initRouter() (e error) {
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	gin.SetMode(a.cfg.Mode)
	a.router = gin.Default()

	store := memstore.NewStore([]byte("sdfasdfasdfasdfkjklkl dfkskasdfasdfkkasdjfaskdjfas;lkdfkkdlllll"))
	a.router.Use(sessions.Sessions("sid", store))
	a.router.Use(auth.SetFlags)

	a.router.SetFuncMap(template.FuncMap{
		"each": func(n, interval int) bool {
			return n%interval == 0
		},
		"each1": func(n, interval int) bool {
			return (n+1)%interval == 0
		},
	})

	// load templates
	a.router.LoadHTMLGlob(a.cfg.TemplatePath + "**/*.tpl")

	a.router.Static("/images/", a.cfg.ImagePath)
	a.router.Static("/bootstrap/", a.cfg.BootstrapPath)
	a.router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/book/")
	})

	a.srv = &http.Server{
		Addr:    a.cfg.Host,
		Handler: a.router.Handler(),
	}

	return nil
}

func (a *storeApp) initBooks() error {

	br, e := repo.NewBookRepo(a.db)
	if e != nil {
		return fmt.Errorf("error create book repo: %w", e)
	}

	a.bookService, e = service.NewBookService(br)
	if e != nil {
		return fmt.Errorf("error create book service: %w", e)
	}

	bc, e := controller.NewBookController(a.bookService)
	if e != nil {
		return fmt.Errorf("error create book controller: %w", e)
	}

	e = bc.Init(&a.router.RouterGroup)
	if e != nil {
		return fmt.Errorf("error init book controller: %w", e)
	}

	return nil
}

func (a *storeApp) initAuth() error {

	ar, e := repo2.NewAuthRepo(a.db)
	if e != nil {
		return fmt.Errorf("error create auth repo: %w", e)
	}

	as, e := service2.NewAuthService(ar)
	if e != nil {
		return fmt.Errorf("error create auth service: %w", e)
	}

	ac, e := controller2.NewAuthController(as)
	if e != nil {
		return fmt.Errorf("error create auth controller: %w", e)
	}

	e = ac.Init(&a.router.RouterGroup)
	if e != nil {
		return fmt.Errorf("error init auth controller: %w", e)
	}

	return nil
}

func (a *storeApp) initAdmin() error {

	ac, e := controller3.NewAdminController(a.bookService)
	if e != nil {
		return fmt.Errorf("error create admin controller: %w", e)
	}

	e = ac.Init(&a.router.RouterGroup)
	if e != nil {
		return fmt.Errorf("error init admin controller: %w", e)
	}

	return nil
}
