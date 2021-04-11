package router

import (
	"github.com/ahojcn/ecloud/ctr/controller"
	"github.com/ahojcn/ecloud/ctr/util"
	"github.com/gin-contrib/cors"
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

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8081", "http://127.0.0.1:8081"}
	config.AllowCredentials = true
	//config.AllowAllOrigins = true
	Router.Use(cors.New(config))

	// 用户相关
	Router.POST("/user", controller.CreateUser)
	Router.GET("/user/:id", controller.GetUserInfoById)
	Router.GET("/user", controller.GetUsersInfoByUsername)

	// 会话控制
	Router.POST("/session", controller.Login)
	Router.DELETE("/session", controller.Logout)

	// 服务树
	Router.POST("/tree", controller.CreateTreeNode)
	Router.GET("/tree", controller.GetTreeNodes)

	// 资源 -- 主机
	Router.POST("/host", controller.CreateHost)
	Router.PUT("/host", controller.UpdateHost)
	Router.DELETE("/host/:id", controller.DeleteHost)
	Router.POST("/host_user", controller.CreateHostUser)
	Router.DELETE("/host_user", controller.DeleteHostUser)

	return Router
}
