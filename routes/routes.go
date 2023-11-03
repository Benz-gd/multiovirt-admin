package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"multiovirt-admin/controllers"
	"multiovirt-admin/logger"
	"multiovirt-admin/middleware"
)

func InitRoutes() *gin.Engine {
	start_mode := viper.GetString("app.start_mode")
	switch start_mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	base := r.Group("/api/base")
	ovirt := r.Group("/api/ovirt")
	tools := r.Group("/api/tools")
	zabbix := r.Group("/api/zabbix")
	//base 基础组件路由
	base.POST("/signup", controllers.SignUpHandler)
	base.POST("/login", controllers.LoginHandler)
	//base.GET("/ping", middleware.JwtAuthMiddleware(), func(c *gin.Context) {
	//	controllers.ResponseSuccess(c, controllers.CodeResponseSuccess)
	//})
	//ovirt 路由
	ovirt.Use(middleware.JwtAuthMiddleware())
	{
		ovirt.POST("/createovirtconf", controllers.CreateOvirtConf)
		ovirt.GET("/listovirtconf", controllers.ListOvirtConf)
		ovirt.GET("/ovirtconfdetail/:aliasname", controllers.GetOvirtConfDetail)
	}

	tools.Use(middleware.JwtAuthMiddleware())
	{
		tools.GET("/ping", controllers.Ping)
	}
	zabbix.Use(middleware.JwtAuthMiddleware())
	{
		zabbix.GET("/listhostgroup", controllers.ListHostGroup)
	}
	r.NoRoute(func(c *gin.Context) {
		controllers.ResponseError(c, controllers.CodeNoRoute)
	})

	return r
}
