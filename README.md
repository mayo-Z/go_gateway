目录

- [Go 微服务网关项目介绍](#go-微服务网关项目介绍)
  - [预览图](#预览图)
  - [数据库E-R图](#数据库e-r图)
  - [项目流程图](#项目流程图)
- [实现功能](#实现功能)
  - [后台管理功能](#后台管理功能)
  - [反向代理功能](#反向代理功能)
- [代码帮助](#代码帮助)
  - [运行后端项目](#运行后端项目)
  - [运行前端项目](#运行前端项目)
  - [运行下游服务器](#运行下游服务器)
  - [代码部署](#代码部署)
    - [实体机部署](#实体机部署)
      - [1、每个项目独立部署](#1每个项目独立部署)
      - [2、前后端合并部署](#2前后端合并部署)
    - [docker部署](#docker部署)
    - [k8s部署(linux)](#k8s部署linux)
- [代码注意事项](#代码注意事项)
  - [不足与bug](#不足与bug)
  - [未来可扩展功能](#未来可扩展功能)

# Go 微服务网关项目介绍

这是一个前后端分离的微服务网关服务，主要实现了http、tcp、grpc的反向代理功能，大家可以下载、运行、测试、修改。

项目架构上使用了微服务架构，中间件模块化开发，以及Restful风格的接口设计，使用swagger生成可视化的API接口页面。

项目使用了中间件模块化开发，实现了翻译器、验证器、错误统一处理等功能。

网关实现功能还有：(JWT)权限认证功能、黑白名单功能、(header头)请求头转换功能、流量控制功能、基于主动探测下游节点的负载均衡功能。

技术栈：Vue2.0+Element+Go+gRPC+JWT+MySQL+Redis

后端项目：https://github.com/uptocorrupt/go_gateway

前端项目：https://github.com/uptocorrupt/go_gateway_view

下游测试服务器：https://github.com/uptocorrupt/gateway_server

项目的预览地址：http://www.gogateway.cn/

联系邮箱：hhd5050@foxmail.com

## 预览图

[![xmVkgf.png](https://s1.ax1x.com/2022/09/28/xmVkgf.png)](https://imgse.com/i/xmVkgf)

[![xmVGKU.png](https://s1.ax1x.com/2022/09/28/xmVGKU.png)](https://imgse.com/i/xmVGKU)

[![xmV08x.png](https://s1.ax1x.com/2022/09/28/xmV08x.png)](https://imgse.com/i/xmV08x)

[![xmV2ad.png](https://s1.ax1x.com/2022/09/28/xmV2ad.png)](https://imgse.com/i/xmV2ad)

## 数据库E-R图

[![xmZpsU.png](https://s1.ax1x.com/2022/09/28/xmZpsU.png)](https://imgse.com/i/xmZpsU)

## 项目流程图

[![xmZFo9.png](https://s1.ax1x.com/2022/09/28/xmZFo9.png)](https://imgse.com/i/xmZFo9)

# 实现功能

##  后台管理功能

管理员：登录、退出、修改密码

大盘：流量统计展示、全站服务类型占比、四维度指标

租户：管理列表、租户信息curd、流量统计

服务：服务列表、流量统计、服务信息crud

## 反向代理功能

http、tcp、grpc反向代理

JWT权限认证、黑白名单、header头转换、流量控制、负载均衡之主动探测

# 代码帮助

以下代码均在windows下运行，如在非windows环境下则会另行标注，其他环境请自行修改相关代码

本后端代码已集成前端代码，如仅需要后端代码，可删除dist文件夹，并修改后端文件: router/route.go

```
// router.Static("/dist", "./dist")
```

## 运行后端项目

- 首先git clone 后端项目
```
git clone https://github.com/uptocorrupt/go_gateway.git
```
- 确保本地环境安装了Go 1.12+版本

```
go version
```

- 下载类库依赖

```
go env -w GO111MODULE=on 
go env -w GOPROXY=https://goproxy.cn
cd go_gateway
go mod tidy
```

- 创建 db 并导入数据

```
mysql -u root -p -e "CREATE DATABASE go_gateway"
mysql -u root -p go_gateway < go_gateway.sql
```

- 调整 配置文件

修改 ./conf/dev文件夹中的文件为自己的环境配置。


- 运行面板、代理服务

运行管理面板配合前端项目 - 达成服务管理功能
```
go run main.go -config=./conf/dev/ -endpoint dashboard
```

运行代理服务
```
go run main.go -config=./conf/dev/ -endpoint server
```

## 运行前端项目

- 首先git clone 前端项目 

```
git clone https://github.com/uptocorrupt/go_gateway_view.git
```

- 确保本地环境安装了nodejs

```
node -v
```

- 安装node依赖包

```
cd go_gateway_view
npm install
```
​	或者使用国内代理
```
npm install -g cnpm --registry=https://registry.npm.taobao.org
cnpm install
```

- 运行前端项目

```
npm run dev
```
## 运行下游服务器 
- git clone

```
git clone https://github.com/uptocorrupt/gateway_server.git
```
- 确保本地环境安装了Go 1.12+版本
```
go version
```

- 下载类库依赖

```
go env -w GO111MODULE=on 
go env -w GOPROXY=https://goproxy.cn
cd go_gateway
go mod tidy
```
- 修改对应服务器的下游ip地址和端口
- 运行

```
go run main.go
```


## 代码部署

### 实体机部署

#### 1、每个项目独立部署

- 前端项目一个端口
- 接口项目一个端口
- 前端设置:
.env.production
```
VUE_APP_BASE_API = 'prod-api '
```
vue.config.js
```
publicPath: '/'
```
- 使用nginx将后端接口设置到跟前端同域下访问
```
    server {
        listen       80
        server_name  localhost;
        root C:\ gateway\go_gateway_view\dist;

        location /prod-api/ {
            proxy_pass http://127.0.0.1:8880/;
        }
    }
```

- 前端打包

```
  npm run build:prod
```

- 代理服务器独立部署

- 启动后端项目

  或者可参考文件里的onekeyupdate.sh和onekeyupdate.cmd制作一键启动脚本

  在linux下启动: vim onekeyupdate.sh

#### 2、前后端合并部署
- 前端打包的dist放到后端同一项目中
- 后端设置: router/route.go
```
router.Static("/dist", "./dist")
```
- 前端设置:
.env.production
```
VUE_APP_BASE_API = ' '
```
vue.config.js
```
publicPath: '/dist'
```
- 启动后端项目
- 在浏览器中打开：http://127.0.0.1:8880/dist/

### docker部署

- 创建docker文件 vim dockerfile-dashboard
- 创建docker镜像

其中"-f"后接参数为自己编写的dockerfile文件，"-t"后接参数为生成的docker镜像名

```
docker build -f dockerfile-dashboard -t dockerfile_dashboard .
docker build -f dockerfile-server -t dockerfile_server .
```
- 运行测试docker镜像:
```
docker run -it --rm --name dockerfile_dashboard -p 8880:8880 dockerfile_dashboard
docker run -it --rm --name dockerfile_server -p 8081:8081 -p 4433:4433 dockerfile_server
```

### k8s部署(linux)

- 创建docker镜像

- 创建交叉编译脚本，解决build太慢问题 
```
vim docker_build.sh
sh ./docker_build.sh
```
- 编写服务编排文件
```
vim k8s_dashboard.yaml
vim k8s_server.yaml
```
- 启动服务
```
kubectl apply -f k8s_dashboard.yaml
kubectl apply -f k8s_server.yaml
```
- 查看所有部署
```
kubectl get all
```

# 代码注意事项

## 不足与bug

本代码目前仍存在较多不足，且存在少量bug，请谨慎使用！

- 在grpc反向代理中间件中，修改metadata头内容失败，导致应用无法修改metadata头，以及无法统计租户在用户使用grpc反向代理功能时的日请求数和QPS数据
- 部分代码引用了不维护的lib包，可能导致未来某些代码运行出现错误
   - <u>github.com/e421083458/golang_common/lib</u>
   - <u>github.com/e421083458/gorm</u>
   - <u>dgrijalva/jwt-go</u> =><u>golang-jwt/jwt</u>


- 在更新下游节点时需重启服务器

## 未来可扩展功能

- 配置热更新，在更新下游节点时无需重启服务器
- 增加细粒度的限流比如基于小时、分钟级别的
- 分布式限流，现为单机版本
- 服务发现之处zk、ectd，目前该服务使用主动探测是服务发现
- 设置权限验证中的access_token自动更新
- gin-jwt实现续签功能
