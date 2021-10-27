package routes

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Start() error
}

type _BasicRouter struct {
	server    *gin.Engine
	groupBase *gin.RouterGroup
}

func Default() Router {
	return &_BasicRouter{}
}

func (router *_BasicRouter) setup() {
	server := gin.Default()
	router.server = server
	server.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
}

func (router *_BasicRouter) loadBase() {
	server := router.server
	router.groupBase = server.Group("/api/v1")
}

func (router *_BasicRouter) Start() error {
	router.setup()
	router.loadBase()
	router.loadAppRoutes()
	return router.server.Run()
}
