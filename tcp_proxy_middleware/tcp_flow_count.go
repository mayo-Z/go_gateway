package tcp_proxy_middleware

import (
	"fmt"
	"go_gateway/dao"
	"go_gateway/public"
	"time"
)

func TCPFlowCountMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		serverInterface := c.Get("service")
		if serverInterface == nil {
			c.conn.Write([]byte("get service empty"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		//全站、服务、租户
		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			c.conn.Write([]byte(err.Error()))
			c.Abort()
			return
		}
		totalCounter.Increase()
		dayCount, _ := totalCounter.GetDayData(time.Now())
		fmt.Printf("totalCounter qps:%v,dayCount:%v\n",
			totalCounter.QPS, dayCount)

		serviceCounter, err := public.FlowCounterHandler.GetCounter(
			public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			c.conn.Write([]byte(err.Error()))
			c.Abort()
			return
		}
		serviceCounter.Increase()
		dayServiceCount, _ := serviceCounter.GetDayData(time.Now())
		fmt.Printf("serviceCounter qps:%v,dayServiceCount:%v\n",
			serviceCounter.QPS, dayServiceCount)

		c.Next()
	}
}
