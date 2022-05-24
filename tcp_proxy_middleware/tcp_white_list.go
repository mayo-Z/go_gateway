package tcp_proxy_middleware

import (
	"fmt"
	"go_gateway/dao"
	"go_gateway/public"
	"strings"
)

//http://127.0.0.1:20004/test_http_string3/abbb
//http://127.0.0.1:20004/abbb
func TCPWhiteListMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)
		ipList := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			ipList = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}

		//ip:port
		splits := strings.Split(c.conn.RemoteAddr().String(), ":")
		clientIP := ""
		if len(splits) == 2 {
			clientIP = splits[0]
		}

		if serviceDetail.AccessControl.OpenAuth == 1 && len(ipList) > 0 {
			if !public.InStringSlice(ipList, clientIP) {
				c.conn.Write([]byte(fmt.Sprintf("%s not in white ip list",
					clientIP)))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
