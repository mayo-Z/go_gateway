package http_proxy_router

import (
	"github.com/gin-gonic/gin"
	"go_gateway/controller"
	"go_gateway/http_proxy_middleware"
	"go_gateway/middleware"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	//router := gin.New()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		return
	})

	oauth := router.Group("/oauth")
	oauth.Use(middleware.TranslationMiddleware())
	{
		controller.OAuthRegister(oauth)
	}

	//中间件顺序不能弄错！！！
	//匹配接入方式、header头转换、最后反向代理
	router.Use(http_proxy_middleware.HTTPAccessModeMiddleware(),

		http_proxy_middleware.HTTPFlowCountMiddleware(),
		http_proxy_middleware.HTTPFlowLimitMiddleware(),

		http_proxy_middleware.HTTPJwtAuthTokenMiddleware(),
		http_proxy_middleware.HTTPJwtFlowCountMiddleware(),
		http_proxy_middleware.HTTPJwtFlowLimitMiddleware(),

		http_proxy_middleware.HTTPWhiteListMiddleware(),
		http_proxy_middleware.HTTPBlackListMiddleware(),

		http_proxy_middleware.HTTPHeaderTransferMiddleware(),
		http_proxy_middleware.HTTPStripUriMiddleware(),
		http_proxy_middleware.HTTPUrlRewriteMiddleware(),

		http_proxy_middleware.HTTPReverseProxyMiddleware(),
	)

	return router
}
