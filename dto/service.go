package dto

import (
	"github.com/gin-gonic/gin"
	"go_gateway/public"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`                      //关键词
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`        //页数
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"20" validate:"required"` //每页条数
}

type ServiceDetailInput struct {
	ID int64 `json:"id" form:"id" comment:"服务ID" example:"56" validate:"required"` //关键词
}

func (param *ServiceListInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

func (param *ServiceDetailInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

type ServiceListItemOutput struct {
	ID          int64  `json:"id" form:"id" `                    //id
	ServiceName string `json:"service_name" form:"service_name"` //服务名称
	ServiceDesc string `json:"service_desc" form:"service_desc"` //服务描述
	LoadType    int    `json:"load_type" form:"load_type"`       //类型
	ServiceAddr string `json:"service_addr" form:"service_addr"` //服务地址
	Qps         int64  `json:"qps" form:"qps"`                   //qps
	Qpd         int64  `json:"qpd" form:"qpd"`                   //qpd
	TotalNode   int    `json:"total_node" form:"total_node"`     //节点数
}

type ServiceListOutput struct {
	Total int64                   `json:"total" form:"total" comment:"总数" example:"" validate:""` //总数
	List  []ServiceListItemOutput `json:"list" form:"list" comment:"列表" example:"" validate:""`   //列表
}

type ServiceStatOutput struct {
	Today     []int64 `json:"today" form:"today" comment:"今日流量" example:"" validate:""`         //总数
	Yesterday []int64 `json:"yesterday" form:"yesterday" comment:"昨日流量" example:"" validate:""` //列表
}

type ServiceAddHTTPInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,max=255,min=1"`

	RuleType       int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"" validate:"max=1,min=0"`
	Rule           string `json:"rule" form:"rule" comment:"接入路径: 域名或者前缀" example:"" validate:"required,valid_rule"`
	NeedHttps      int    `json:"need_https" form:"need_https" comment:"支持https" example:"" validate:"max=1,min=0"`
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"启用websocket 1=启用" example:"" validate:"max=1,min=0"`
	NeedStripUri   int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri 1=启用" example:"" validate:"max=1,min=0"`
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"URL重写" example:"" validate:"valid_url_rewrite"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"Header转换" example:"" validate:"valid_header_transfor"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限 1=开启" example:"" validate:"max=1,min=0"`
	BlackList         string `json:"black_list" form:"black_list" comment:"IP黑名单" example:"" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"IP白名单" example:"" validate:"valid_ip_list"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流" example:"" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式 round/weight_round/random/ip_hash" example:"" validate:"max=3,min=0"`
	IpList                 string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"" validate:"required,valid_ip_port_list"`
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weight_list"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"下游建立连接超时, 单位s" example:"" validate:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"下游获取header超时, 单位s" example:"" validate:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"下游链接最大空闲时间, 单位s" example:"" validate:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"下游最大空闲链接数" example:"" validate:"min=0"`
}

func (param *ServiceAddHTTPInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

type ServiceUpdateHTTPInput struct {
	ID          int64  `json:"id" form:"id" comment:"服务ID" example:"62" validate:"required,min=1"`
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"test_http_service_indb" validate:"required,valid_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"test_http_service_indb" validate:"required,max=255,min=1"`

	RuleType       int    `json:"rule_type" form:"rule_type" comment:"接入类型" example:"" validate:"max=1,min=0"`
	Rule           string `json:"rule" form:"rule" comment:"接入路径: 域名或者前缀" example:"/test_http_service_indb" validate:"required,valid_rule"`
	NeedHttps      int    `json:"need_https" form:"need_https" comment:"支持https" example:"" validate:"max=1,min=0"`
	NeedWebsocket  int    `json:"need_websocket" form:"need_websocket" comment:"启用websocket 1=启用" example:"" validate:"max=1,min=0"`
	NeedStripUri   int    `json:"need_strip_uri" form:"need_strip_uri" comment:"启用strip_uri 1=启用" example:"" validate:"max=1,min=0"`
	UrlRewrite     string `json:"url_rewrite" form:"url_rewrite" comment:"URL重写" example:"" validate:"valid_url_rewrite"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"Header转换" example:"" validate:"valid_header_transfor"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限 1=开启" example:"" validate:"max=1,min=0"`
	BlackList         string `json:"black_list" form:"black_list" comment:"IP黑名单" example:"" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"IP白名单" example:"" validate:"valid_ip_list"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流" example:"" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`

	RoundType              int    `json:"round_type" form:"round_type" comment:"轮询方式 round/weight_round/random/ip_hash" example:"" validate:"max=3,min=0"`
	IpList                 string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"127.0.0.1:80" validate:"required,valid_ip_port_list"`
	WeightList             string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"50" validate:"required,valid_weight_list"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" comment:"下游建立连接超时, 单位s" example:"" validate:"min=0"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" comment:"下游获取header超时, 单位s" example:"" validate:"min=0"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" comment:"下游链接最大空闲时间, 单位s" example:"" validate:"min=0"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" comment:"下游最大空闲链接数" example:"" validate:"min=0"`
}

func (param *ServiceUpdateHTTPInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

type ServiceAddTcpInput struct {
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,max=255,min=1"`
	Port        int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" example:"" validate:"min=8001,max=8999"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限 1=开启" example:"" validate:"max=1,min=0"`
	BlackList         string `json:"black_list" form:"black_list" comment:"IP黑名单" example:"" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"IP白名单" example:"" validate:"valid_ip_list"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流" example:"" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`

	RoundType  int    `json:"round_type" form:"round_type" comment:"轮询方式 round/weight_round/random/ip_hash" example:"" validate:"max=3,min=0"`
	IpList     string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"" validate:"required,valid_ip_port_list"`
	WeightList string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weight_list"`
}

func (param *ServiceAddTcpInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

type ServiceUpdateTcpInput struct {
	ID          int64  `json:"id" form:"id" comment:"服务ID" example:"62" validate:"required,min=1"`
	ServiceName string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,max=255,min=1"`
	Port        int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" example:"" validate:"min=8001,max=8999"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限 1=开启" example:"" validate:"max=1,min=0"`
	BlackList         string `json:"black_list" form:"black_list" comment:"IP黑名单" example:"" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"IP白名单" example:"" validate:"valid_ip_list"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流" example:"" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`

	RoundType  int    `json:"round_type" form:"round_type" comment:"轮询方式 round/weight_round/random/ip_hash" example:"" validate:"max=3,min=0"`
	IpList     string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"" validate:"required,valid_ip_port_list"`
	WeightList string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weight_list"`
}

func (param *ServiceUpdateTcpInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

type ServiceAddGrpcInput struct {
	ServiceName    string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"`
	ServiceDesc    string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,max=255,min=1"`
	Port           int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" example:"" validate:"min=8001,max=8999"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" example:"" validate:"valid_header_transfor"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限 1=开启" example:"" validate:"max=1,min=0"`
	BlackList         string `json:"black_list" form:"black_list" comment:"IP黑名单" example:"" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"IP白名单" example:"" validate:"valid_ip_list"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流" example:"" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`

	RoundType  int    `json:"round_type" form:"round_type" comment:"轮询方式 round/weight_round/random/ip_hash" example:"" validate:"max=3,min=0"`
	IpList     string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"" validate:"required,valid_ip_port_list"`
	WeightList string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weight_list"`
}

func (param *ServiceAddGrpcInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}

type ServiceUpdateGrpcInput struct {
	ID             int64  `json:"id" form:"id" comment:"服务ID" example:"62" validate:"required,min=1"`
	ServiceName    string `json:"service_name" form:"service_name" comment:"服务名" example:"" validate:"required,valid_service_name"`
	ServiceDesc    string `json:"service_desc" form:"service_desc" comment:"服务描述" example:"" validate:"required,max=255,min=1"`
	Port           int    `json:"port" form:"port" comment:"端口，需要设置8001-8999范围内" example:"" validate:"min=8001,max=8999"`
	HeaderTransfor string `json:"header_transfor" form:"header_transfor" comment:"metadata转换" example:"" validate:"valid_header_transfor"`

	OpenAuth          int    `json:"open_auth" form:"open_auth" comment:"是否开启权限 1=开启" example:"" validate:"max=1,min=0"`
	BlackList         string `json:"black_list" form:"black_list" comment:"IP黑名单" example:"" validate:"valid_ip_list"`
	WhiteList         string `json:"white_list" form:"white_list" comment:"IP白名单" example:"" validate:"valid_ip_list"`
	ClientIPFlowLimit int    `json:"clientip_flow_limit" form:"clientip_flow_limit" comment:"客户端ip限流" example:"" validate:"min=0"`
	ServiceFlowLimit  int    `json:"service_flow_limit" form:"service_flow_limit" comment:"服务端限流" example:"" validate:"min=0"`
	WhiteHostName     string `json:"white_host_name" form:"white_host_name" comment:"白名单主机，以逗号间隔" validate:"valid_ip_list"`
	ForbidList        string `json:"forbid_list" form:"forbid_list" comment:"禁用IP列表" validate:"valid_ip_list"`

	RoundType  int    `json:"round_type" form:"round_type" comment:"轮询方式 round/weight_round/random/ip_hash" example:"" validate:"max=3,min=0"`
	IpList     string `json:"ip_list" form:"ip_list" comment:"ip列表" example:"" validate:"required,valid_ip_port_list"`
	WeightList string `json:"weight_list" form:"weight_list" comment:"权重列表" example:"" validate:"required,valid_weight_list"`
}

func (param *ServiceUpdateGrpcInput) BindValidParam(ctx *gin.Context) error {
	return public.DefaultGetValidParams(ctx, param)
}
