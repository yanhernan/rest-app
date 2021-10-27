package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getApp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func (router *_BasicRouter) loadAppRoutes() {
	group := router.groupBase
	appGroup := group.Group("/apps")
	appGroup.GET("/", getApp)
}
