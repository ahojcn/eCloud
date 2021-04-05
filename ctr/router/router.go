package router

import (
	"github.com/ahojcn/ecloud/ctr/controller"
	"github.com/ahojcn/ecloud/ctr/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Router *gin.Engine

func init() {
	Router = gin.Default()
}

func SetupRouter() *gin.Engine {
	store := memstore.NewStore([]byte(util.Config.Section("system").Key("session_secret").String()))
	Router.Use(sessions.Sessions(util.Config.Section("system").Key("session_name").String(), store))

	Router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "404")
	})

	Router.POST("/user", controller.CreateUser)
	Router.GET("/user/:id", controller.GetUserInfoById)

	return Router
}
