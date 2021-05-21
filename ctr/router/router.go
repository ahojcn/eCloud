package router

import (
	"net/http"

	"github.com/ahojcn/ecloud/ctr/controller"
	"github.com/ahojcn/ecloud/ctr/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-gonic/gin"
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
	config.AllowOrigins = []string{"http://localhost:8080", "http://127.0.0.1:8080"}
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
	Router.GET("/session", controller.IsLogin)

	// 服务树
	Router.POST("/tree", controller.CreateTreeNode)
	Router.GET("/tree", controller.GetTreeNodes)
	Router.DELETE("/tree", controller.DeleteTreeNode)
	Router.POST("/user_tree", controller.CreateUserTree)
	Router.DELETE("/user_tree", controller.DeleteUserTree)

	// 资源 -- 主机
	Router.GET("/host", controller.GetHostInfo)
	Router.POST("/host", controller.CreateHost)
	Router.PUT("/host", controller.UpdateHost)
	Router.DELETE("/host/:id", controller.DeleteHost)
	Router.POST("/host_user", controller.CreateHostUser)
	Router.DELETE("/host_user", controller.DeleteHostUser)
	Router.POST("/command", controller.RunCommand)

	// 监控
	Router.POST("/metrics", controller.MonitorWriteMetrics)
	Router.GET("/metrics", controller.MonitorQueryMetrics)

	Router.POST("/router", controller.RouterMonitorMetricsWrite)
	Router.GET("/router/metrics", controller.RouterMonitorMetricsGet)
	Router.GET("/router/query", controller.RouterMonitorMetricsQuery)

	// ICode
	Router.GET("/icode", controller.GetICodeList)
	Router.POST("/icode", controller.CreateICode)
	Router.DELETE("/icode", controller.DeleteICode)

	// Router
	Router.POST("/m_router", controller.MarkHostAsRouter)
	Router.POST("/m_router/redo", controller.RouterRedo)
	Router.GET("/m_router", controller.RouterList)
	Router.GET("/m_router/status", controller.RouterStatus)
	Router.GET("/m_router/nginx/config", controller.NginxConfig)

	// PipeLine
	Router.GET("/pipeline/list", controller.PipeLineList)
	Router.POST("/pipeline/create", controller.PipeLineCreate)
	Router.POST("/pipeline/run", controller.PipeLineRun)
	Router.GET("/pipeline/status", controller.PipeLineStatus)
	Router.GET("/pipeline/reset", controller.PipeLineReset)
	Router.GET("/pipeline/delete", controller.PipeLineDelete)

	// Cluster
	Router.GET("/cluster/one", controller.ClusterRetrieve)
	Router.GET("/cluster/delete", controller.ClusterDelete)
	Router.POST("/cluster/create", controller.ClusterCreate)
	Router.GET("/cluster/list", controller.ClusterList)

	// Service
	Router.POST("/service/create", controller.ServiceCreate)
	Router.GET("/service/one", controller.ServiceGet)
	Router.GET("/service/delete", controller.ServiceDelete)

	return Router
}
