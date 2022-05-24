package http_proxy_middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gateway/dao"
	"go_gateway/middleware"
	"go_gateway/public"
	"strings"
)

//http://127.0.0.1:20004/test_http_string3/abbb
//http://127.0.0.1:20004/abbb
func HTTPStripUriMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		if serviceDetail.HttpRule.RuleType == public.HTTPRuleTypePrefixURL &&
			serviceDetail.HttpRule.NeedStripUri == 1 {
			fmt.Println("c.Request.URL.Path", c.Request.URL.Path)
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path,
				serviceDetail.HttpRule.Rule, "", 1)
			fmt.Println("c.Request.URL.Path", c.Request.URL.Path)

		}

		c.Next()
	}
}
