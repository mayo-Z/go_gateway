package controller

import (
	"errors"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"go_gateway/dao"
	"go_gateway/dto"
	"go_gateway/middleware"
	"go_gateway/public"
	"strings"
	"time"
)

type ServiceController struct{}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
	group.GET("/service_delete", service.ServiceDelete)
	group.GET("/service_detail", service.ServiceDetail)
	group.GET("/service_stat", service.ServiceStat)

	group.POST("/service_add_http", service.ServiceAddHTTP)
	group.POST("/service_update_http", service.ServiceUpdateHTTP)
	group.POST("/service_add_tcp", service.ServiceAddTcp)
	group.POST("/service_update_tcp", service.ServiceUpdateTcp)
	group.POST("/service_add_grpc", service.ServiceAddGrpc)
	group.POST("/service_update_grpc", service.ServiceUpdateGrpc)
}

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string false "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(ctx *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(ctx, tx, params)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	outList := []dto.ServiceListItemOutput{}
	for _, listItem := range list {
		serviceDetail, err := listItem.ServiceDetail(ctx, tx, &listItem)
		if err != nil {
			middleware.ResponseError(ctx, 2003, err)
			return
		}

		//1.http后缀介入clusterIP+clusterPort+path
		//2.http域名接入domain
		//3.tcp、grpc接入clusterIP+servicePort
		serviceAddr := "unknown"
		clusterIP := lib.GetStringConf("base.cluster.cluster_ip")
		clusterPort := lib.GetStringConf("base.cluster.cluster_port")
		clusterSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")

		if serviceDetail.Info.LoadType == public.LoadTypeHTTP &&
			serviceDetail.HttpRule.RuleType == public.HTTPRuleTypePrefixURL && serviceDetail.HttpRule.NeedHttps == 1 {
			serviceAddr = fmt.Sprintf("%s:%s%s", clusterIP, clusterSSLPort, serviceDetail.HttpRule.Rule)
		}
		if serviceDetail.Info.LoadType == public.LoadTypeHTTP &&
			serviceDetail.HttpRule.RuleType == public.HTTPRuleTypePrefixURL && serviceDetail.HttpRule.NeedHttps == 0 {
			serviceAddr = fmt.Sprintf("%s:%s%s", clusterIP, clusterPort, serviceDetail.HttpRule.Rule)
		}
		if serviceDetail.Info.LoadType == public.LoadTypeHTTP &&
			serviceDetail.HttpRule.RuleType == public.HTTPRuleTypeDomain {
			serviceAddr = serviceDetail.HttpRule.Rule
		}

		if serviceDetail.Info.LoadType == public.LoadTypeTCP {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.TCPRule.Port)
		}
		if serviceDetail.Info.LoadType == public.LoadTypeGRPC {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.GRPCRule.Port)
		}
		ipList := serviceDetail.LoadBalance.GetIPListByModel()
		counter, err := public.FlowCounterHandler.GetCounter(
			public.FlowServicePrefix + listItem.ServiceName)
		if err != nil {
			middleware.ResponseError(ctx, 2004, err)
			return
		}
		outItem := dto.ServiceListItemOutput{
			ID:          listItem.ID,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			ServiceAddr: serviceAddr,
			Qps:         counter.QPS,
			Qpd:         counter.TotalCount,
			TotalNode:   len(ipList),
			LoadType:    listItem.LoadType,
		}

		outList = append(outList, outItem)
	}
	out := &dto.ServiceListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(ctx, out)
}

// ServiceDetail godoc
// @Summary 服务详情
// @Description 服务详情
// @Tags 服务管理
// @ID /service/service_detail
// @Accept  json
// @Produce  json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=dao.ServiceDetail} "success"
// @Router /service/service_detail [get]
func (service *ServiceController) ServiceDetail(ctx *gin.Context) {
	params := &dto.ServiceDetailInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}

	//读取基本信息
	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceInfo, err = serviceInfo.Find(ctx, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(ctx, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(ctx, 2003, err)
		return
	}
	middleware.ResponseSuccess(ctx, serviceDetail)
}

// ServiceStat godoc
// @Summary 服务统计
// @Description 服务统计
// @Tags 服务管理
// @ID /service/service_stat
// @Accept  json
// @Produce  json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=dto.ServiceStatOutput} "success"
// @Router /service/service_stat [get]
func (service *ServiceController) ServiceStat(ctx *gin.Context) {
	params := &dto.ServiceDetailInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}

	//读取基本信息
	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceDetail, err := serviceInfo.ServiceDetail(ctx, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(ctx, 2003, err)
		return
	}

	counter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
	if err != nil {
		middleware.ResponseError(ctx, 2004, err)
		return
	}

	todayList := []int64{}
	currentTime := time.Now()
	for i := 0; i <= currentTime.Hour(); i++ {
		dateTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(),
			i, 0, 0, 0, lib.TimeLocation)
		hourData, _ := counter.GetHourData(dateTime)
		todayList = append(todayList, hourData)
	}
	yesterdayList := []int64{}
	yesterTime := currentTime.Add(-1 * time.Duration(time.Hour*24))
	for i := 0; i <= 23; i++ {
		dateTime := time.Date(yesterTime.Year(), yesterTime.Month(), yesterTime.Day(),
			i, 0, 0, 0, lib.TimeLocation)
		hourData, _ := counter.GetHourData(dateTime)
		yesterdayList = append(yesterdayList, hourData)
	}

	middleware.ResponseSuccess(ctx, &dto.ServiceStatOutput{
		Today:     todayList,
		Yesterday: yesterdayList,
	})
}

// ServiceDelete godoc
// @Summary 服务删除
// @Description 服务删除
// @Tags 服务管理
// @ID /service/service_delete
// @Accept  json
// @Produce  json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_delete [get]
func (service *ServiceController) ServiceDelete(ctx *gin.Context) {
	params := &dto.ServiceDetailInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2001, err)
		return
	}
	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceInfo, err = serviceInfo.Find(ctx, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	serviceInfo.IsDelete = 1
	if err = serviceInfo.Save(ctx, tx); err != nil {
		middleware.ResponseError(ctx, 2003, err)
		return
	}
	middleware.ResponseSuccess(ctx, "")
}

// ServiceAddHTTP godoc
// @Summary 添加HTTP服务
// @Description 添加HTTP服务
// @Tags 服务管理
// @ID /service/service_add_http
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceAddHTTPInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_http [post]
func (service *ServiceController) ServiceAddHTTP(ctx *gin.Context) {
	params := &dto.ServiceAddHTTPInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}

	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(ctx, 2001, errors.New("IP列表与权重列表数量不一致"))
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	tx = tx.Begin()
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	_, err = serviceInfo.Find(ctx, tx, serviceInfo)
	if err == nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2003, errors.New("服务已存在"))
		return
	}

	httpUrl := &dao.HttpRule{RuleType: params.RuleType, Rule: params.Rule}
	_, err = httpUrl.Find(ctx, tx, httpUrl)
	if err == nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2004, errors.New("服务接入前缀或域名已存在"))
		return
	}

	serviceModel := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}
	if err := serviceModel.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2005, err)
		return
	}

	httpRule := &dao.HttpRule{
		ServiceID:      serviceModel.ID,
		RuleType:       params.RuleType,
		Rule:           params.Rule,
		NeedHttps:      params.NeedHttps,
		NeedWebsocket:  params.NeedWebsocket,
		NeedStripUri:   params.NeedStripUri,
		UrlRewrite:     params.UrlRewrite,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err := httpRule.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2006, err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         serviceModel.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
	}
	if err := accessControl.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2007, err)
		return
	}

	loadbalance := &dao.LoadBalance{
		ServiceID:              serviceModel.ID,
		RoundType:              params.RoundType,
		IpList:                 params.IpList,
		WeightList:             params.WeightList,
		UpstreamConnectTimeout: params.UpstreamConnectTimeout,
		UpstreamHeaderTimeout:  params.UpstreamHeaderTimeout,
		UpstreamIdleTimeout:    params.UpstreamIdleTimeout,
		UpstreamMaxIdle:        params.UpstreamMaxIdle,
	}
	if err := loadbalance.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2008, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(ctx, "")
	return
}

// ServiceUpdateHTTP godoc
// @Summary 更新HTTP服务
// @Description 更新HTTP服务
// @Tags 服务管理
// @ID /service/service_update_http
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceUpdateHTTPInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_update_http [post]
func (service *ServiceController) ServiceUpdateHTTP(ctx *gin.Context) {
	params := &dto.ServiceUpdateHTTPInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}

	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(ctx, 2001, errors.New("IP列表与权重列表数量不一致"))
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(ctx, 2002, err)
		return
	}
	tx = tx.Begin()
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err = serviceInfo.Find(ctx, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2003, errors.New("服务不存在"))
		return
	}
	serviceDetail, err := serviceInfo.ServiceDetail(ctx, tx, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2004, errors.New("服务不存在"))
		return
	}

	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2005, err)
		return
	}

	//serviceModel.ID
	httpRule := serviceDetail.HttpRule
	httpRule.NeedHttps = params.NeedHttps
	httpRule.NeedStripUri = params.NeedStripUri
	httpRule.NeedWebsocket = params.NeedWebsocket
	httpRule.UrlRewrite = params.UrlRewrite
	httpRule.HeaderTransfor = params.HeaderTransfor
	if err := httpRule.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2006, err)
		return
	}

	accessControl := serviceDetail.AccessControl
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2007, err)
		return
	}

	loadbalance := serviceDetail.LoadBalance
	loadbalance.RoundType = params.RoundType
	loadbalance.IpList = params.IpList
	loadbalance.WeightList = params.WeightList
	loadbalance.UpstreamConnectTimeout = params.UpstreamConnectTimeout
	loadbalance.UpstreamHeaderTimeout = params.UpstreamHeaderTimeout
	loadbalance.UpstreamIdleTimeout = params.UpstreamIdleTimeout
	loadbalance.UpstreamMaxIdle = params.UpstreamMaxIdle
	if err := loadbalance.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2008, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(ctx, "")
	return
}

// ServiceAddTcp godoc
// @Summary 添加tcp服务
// @Description 添加tcp服务
// @Tags 服务管理
// @ID /service/service_add_tcp
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceAddTcpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_tcp [post]
func (service *ServiceController) ServiceAddTcp(ctx *gin.Context) {
	params := &dto.ServiceAddTcpInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}

	//验证service_name是否被占用
	infoSearch := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		IsDelete:    0}

	if _, err := infoSearch.Find(ctx, lib.GORMDefaultPool, infoSearch); err == nil {
		middleware.ResponseError(ctx, 2001, errors.New("服务名被占用,请重新输入"))
		return
	}

	//验证端口是否被占用
	tcpRuleSearch := &dao.TcpRule{Port: params.Port}
	if _, err := tcpRuleSearch.Find(ctx, lib.GORMDefaultPool, tcpRuleSearch); err == nil {
		middleware.ResponseError(ctx, 2002, errors.New("服务端口被占用,请重新输入"))
		return
	}
	grpcRuleSearch := &dao.GrpcRule{Port: params.Port}
	if _, err := grpcRuleSearch.Find(ctx, lib.GORMDefaultPool, grpcRuleSearch); err == nil {
		middleware.ResponseError(ctx, 2003, errors.New("服务端口被占用,请重新输入"))
		return
	}

	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(ctx, 2004, errors.New("IP列表与权重列表数量不一致"))
		return
	}
	tx := lib.GORMDefaultPool.Begin()
	serviceModel := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
		LoadType:    public.LoadTypeTCP,
	}
	if err := serviceModel.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2005, err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         serviceModel.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
		WhiteHostName:     params.WhiteHostName,
	}
	if err := accessControl.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2006, err)
		return
	}

	tcpRule := &dao.TcpRule{
		ServiceID: serviceModel.ID,
		Port:      params.Port,
	}

	if err := tcpRule.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2007, err)
		return
	}

	loadbalance := &dao.LoadBalance{
		ServiceID:  serviceModel.ID,
		RoundType:  params.RoundType,
		IpList:     params.IpList,
		WeightList: params.WeightList,
		ForbidList: params.ForbidList,
	}
	if err := loadbalance.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2008, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(ctx, "")
	return
}

// ServiceUpdateTcp godoc
// @Summary 更新tcp服务
// @Description 更新tcp服务
// @Tags 服务管理
// @ID /service/service_update_tcp
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceUpdateTcpInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_update_tcp [post]
func (service *ServiceController) ServiceUpdateTcp(ctx *gin.Context) {
	params := &dto.ServiceUpdateTcpInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(ctx, 2001, errors.New("IP列表与权重列表数量不一致"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()
	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceDetail, err := serviceInfo.ServiceDetail(ctx, lib.GORMDefaultPool, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2002, errors.New("服务不存在"))
		return
	}
	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc

	if err := info.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2003, err)
		return
	}

	tcpRule := &dao.TcpRule{}
	if serviceDetail.TCPRule != nil {
		tcpRule = serviceDetail.TCPRule
	}
	tcpRule.ServiceID = info.ID
	tcpRule.Port = params.Port
	if err := tcpRule.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2004, err)
		return
	}

	loadBalance := &dao.LoadBalance{}
	if serviceDetail.LoadBalance != nil {
		loadBalance = serviceDetail.LoadBalance
	}
	loadBalance.ServiceID = info.ID
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err := loadBalance.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2005, err)
		return
	}

	accessControl := &dao.AccessControl{}
	if serviceDetail.AccessControl != nil {
		accessControl = serviceDetail.AccessControl
	}
	accessControl.ServiceID = info.ID
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2006, err)
		return
	}

	tx.Commit()
	middleware.ResponseSuccess(ctx, "")
	return
}

// ServiceAddGrpc godoc
// @Summary 添加grpc服务
// @Description 添加grpc服务
// @Tags 服务管理
// @ID /service/service_add_grpc
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceAddGrpcInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_grpc [post]
func (service *ServiceController) ServiceAddGrpc(ctx *gin.Context) {
	params := &dto.ServiceAddGrpcInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}

	//验证service_name是否被占用
	infoSearch := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		IsDelete:    0}

	if _, err := infoSearch.Find(ctx, lib.GORMDefaultPool, infoSearch); err == nil {
		middleware.ResponseError(ctx, 2001, errors.New("服务名被占用,请重新输入"))
		return
	}

	//验证端口是否被占用
	tcpRuleSearch := &dao.TcpRule{Port: params.Port}
	if _, err := tcpRuleSearch.Find(ctx, lib.GORMDefaultPool, tcpRuleSearch); err == nil {
		middleware.ResponseError(ctx, 2002, errors.New("服务端口被占用,请重新输入"))
		return
	}
	grpcRuleSearch := &dao.GrpcRule{Port: params.Port}
	if _, err := grpcRuleSearch.Find(ctx, lib.GORMDefaultPool, grpcRuleSearch); err == nil {
		middleware.ResponseError(ctx, 2003, errors.New("服务端口被占用,请重新输入"))
		return
	}

	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(ctx, 2004, errors.New("IP列表与权重列表数量不一致"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()

	serviceModel := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
		LoadType:    public.LoadTypeGRPC,
	}
	if err := serviceModel.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2005, err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID:         serviceModel.ID,
		OpenAuth:          params.OpenAuth,
		BlackList:         params.BlackList,
		WhiteList:         params.WhiteList,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit:  params.ServiceFlowLimit,
		WhiteHostName:     params.WhiteHostName,
	}
	if err := accessControl.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2006, err)
		return
	}

	grpcRule := &dao.GrpcRule{
		ServiceID:      serviceModel.ID,
		Port:           params.Port,
		HeaderTransfor: params.HeaderTransfor}

	if err := grpcRule.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2007, err)
		return
	}

	loadBalance := &dao.LoadBalance{
		ServiceID:  serviceModel.ID,
		RoundType:  params.RoundType,
		IpList:     params.IpList,
		WeightList: params.WeightList,
		ForbidList: params.ForbidList,
	}
	if err := loadBalance.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2008, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(ctx, "")
	return
}

// ServiceUpdateGrpc godoc
// @Summary 更新grpc服务
// @Description 更新grpc服务
// @Tags 服务管理
// @ID /service/service_update_grpc
// @Accept  json
// @Produce  json
// @Param body body dto.ServiceUpdateGrpcInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_update_grpc [post]
func (service *ServiceController) ServiceUpdateGrpc(ctx *gin.Context) {
	params := &dto.ServiceUpdateGrpcInput{}
	if err := params.BindValidParam(ctx); err != nil {
		middleware.ResponseError(ctx, 2000, err)
		return
	}

	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")) {
		middleware.ResponseError(ctx, 2001, errors.New("IP列表与权重列表数量不一致"))
		return
	}

	tx := lib.GORMDefaultPool.Begin()

	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceDetail, err := serviceInfo.ServiceDetail(ctx, lib.GORMDefaultPool, serviceInfo)
	if err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2002, errors.New("服务不存在"))
		return
	}
	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2003, err)
		return
	}

	grpcRule := &dao.GrpcRule{}
	if serviceDetail.GRPCRule != nil {
		grpcRule = serviceDetail.GRPCRule
	}
	grpcRule.ServiceID = info.ID
	grpcRule.HeaderTransfor = params.HeaderTransfor
	if err := grpcRule.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2004, err)
		return
	}

	loadBalance := &dao.LoadBalance{}
	if serviceDetail.LoadBalance != nil {
		loadBalance = serviceDetail.LoadBalance
	}
	loadBalance.ServiceID = info.ID
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err := loadBalance.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2005, err)
		return
	}

	accessControl := &dao.AccessControl{}
	if serviceDetail.AccessControl != nil {
		accessControl = serviceDetail.AccessControl
	}
	accessControl.ServiceID = info.ID
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlackList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.WhiteHostName = params.WhiteHostName
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err := accessControl.Save(ctx, tx); err != nil {
		tx.Rollback()
		middleware.ResponseError(ctx, 2006, err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(ctx, "")
	return
}
