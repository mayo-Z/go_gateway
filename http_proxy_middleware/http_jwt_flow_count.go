package http_proxy_middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gateway/dao"
	"go_gateway/middleware"
	"go_gateway/public"
)

func HTTPJwtFlowCountMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appInterface, ok := c.Get("app")
		if !ok {
			c.Next()
			return
		}
		appInfo := appInterface.(*dao.App)

		appCounter, err := public.FlowCounterHandler.GetCounter(
			public.FlowAppPrefix + appInfo.AppID)
		if err != nil {
			middleware.ResponseError(c, 2001, err)
			c.Abort()
			return
		}
		appCounter.Increase()
		if appInfo.Qpd > 0 && appCounter.TotalCount > appInfo.Qpd {
			middleware.ResponseError(c, 2002,
				errors.New(fmt.Sprintf("租户日请求量限流 limit:%v current:%v",
					appInfo.Qpd, appCounter.TotalCount)))
			c.Abort()
			return
		}
		c.Next()
	}
}
