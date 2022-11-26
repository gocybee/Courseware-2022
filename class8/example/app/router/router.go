package router

import (
	"github.com/gin-gonic/gin"
	g "main/app/global"
	"main/app/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ZapLogger(g.Logger), middleware.ZapRecovery(g.Logger, true))
	// 使用其他的中间件(跨域,限流...)

	routerGroup := new(Group)

	publicGroup := r.Group("/api")
	{
		routerGroup.InitUserSignRouter(publicGroup)
	}

	privateGroup := r.Group("/api")
	privateGroup.Use(middleware.JWTAuthMiddleware())
	{

	}

	g.Logger.Info("initialize routers successfully!")

	return r
}
