package http_proxy_middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go_gateway/dao"
	"go_gateway/middleware"
	"regexp"
	"strings"
)

//http://127.0.0.1:2004/test_http_string3/abbb
//http://127.0.0.1:2004/abbb
func HTTPUrlRewriteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			middleware.ResponseError(c, 2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		for _, item := range strings.Split(serviceDetail.HttpRule.UrlRewrite, ",") {
			items := strings.Split(item, " ")
			if len(items) != 2 {
				continue
			}

			compile, err := regexp.Compile(items[0])
			if err != nil {
				fmt.Println("regexp.Compile err", err)
				continue
			}
			replacePath := compile.ReplaceAll([]byte(c.Request.URL.Path), []byte(items[1]))
			c.Request.URL.Path = string(replacePath)
		}
		c.Next()
	}
}
