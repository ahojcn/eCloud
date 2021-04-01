package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetupRouter() *gin.Engine {
	Router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "404")
	})

	return Router
}
