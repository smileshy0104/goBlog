package initialize

import (
	"github.com/gin-gonic/gin"
	"goBlog/app/middleware"
	"goBlog/lib/response"
)

// InitRouter 初始化路由
// @title InitRouter
func InitRouter() *gin.Engine {
	r := gin.New()

	//路由404处理
	r.NoRoute(func(c *gin.Context) {
		response.Output(c, 404, "接口不存在", nil)
	})

	//全局中间件 -跨域、
	r.Use(middleware.Cors())
	//异常捕获
	//r.Use(middleware.ExceptionMiddleware)

	//TODO  全局中间件 - 日志 - 异常
	// r.Use(middleware.GinLogger())

	// logger := logs.GetLog()
	// r.Use(zap.Ginzap(logger, time.RFC3339, true))
	// r.Use(ginzap.RecoveryWithZap(logger, true))

	//操作日志中间件
	// r.Use(middleware.LoggerToMysql())

	//InitSysRouter 引入系统基础功能路由
	// InitSysRouter(r)

	return r
}
