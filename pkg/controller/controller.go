package controller

import "github.com/gin-gonic/gin"

type HttpController interface {
	Init(r *gin.RouterGroup) error
}
