package main

import (
	"flag"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"go_gateway/dao"
	"go_gateway/grpc_proxy_router"
	"go_gateway/http_proxy_router"
	"go_gateway/router"
	"go_gateway/tcp_proxy_router"
	"os"
	"os/signal"
	"syscall"
)

//endpoint dashboard后台管理 server代理服务器
//config ./conf/prod/ 对应配置文件夹

var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
	config   = flag.String("config", "", "input config file like ./conf/dev/")
)

func main() {
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *config == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *endpoint == "dashboard" {
		//如果configPath为空，则从命令行中`-config=./conf/prod/`中读取
		//测试用 “go run main.go -config=./conf/dev/ -endpoint=dashboard”
		// 生产环境用“go run main.go -config=./conf/prod/ -endpoint=dashboard”
		lib.InitModule(*config, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		router.HttpServerRun()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		router.HttpServerStop()
	} else {
		//测试用 “go run main.go -config=./conf/dev/ -endpoint=server”
		//生产环境用 “go run main.go -config=./conf/prod/ -endpoint=server”
		lib.InitModule(*config, []string{"base", "mysql", "redis"})
		defer lib.Destroy()
		dao.ServiceManagerHandler.LoadOnce()
		dao.AppManagerHandler.LoadOnce()
		go func() {
			http_proxy_router.HttpServerRun()
		}()
		go func() {
			http_proxy_router.HttpsServerRun()
		}()
		go func() {
			tcp_proxy_router.TcpServerRun()
		}()
		go func() {
			grpc_proxy_router.GrpcServerRun()
		}()

		fmt.Println("start server")

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		tcp_proxy_router.TcpServerStop()
		grpc_proxy_router.GrpcServerStop()
		http_proxy_router.HttpServerStop()
		http_proxy_router.HttpsServerStop()
	}
}
