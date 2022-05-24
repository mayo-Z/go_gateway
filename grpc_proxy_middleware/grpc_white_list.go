package grpc_proxy_middleware

import (
	"errors"
	"fmt"
	"go_gateway/dao"
	"go_gateway/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"strings"
)

//http://127.0.0.1:20004/test_http_string3/abbb
//http://127.0.0.1:20004/abbb
func GrpcWhiteListMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		ipList := []string{}
		if serviceDetail.AccessControl.WhiteList != "" {
			ipList = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}
		peerCtx, ok := peer.FromContext(ss.Context())
		if !ok {
			return errors.New("peer not found with context")
		}
		peerAddr := peerCtx.Addr.String()
		addrPos := strings.LastIndex(peerAddr, ":")
		clientIp := peerAddr[:addrPos]
		if serviceDetail.AccessControl.OpenAuth == 1 && len(ipList) > 0 {
			if !public.InStringSlice(ipList, clientIp) {
				return errors.New(fmt.Sprintf("%s not in white ip list", clientIp))
			}
		}
		if err := handler(srv, ss); err != nil {
			log.Printf("GrpcWhiteListMiddleware failed with error %v\n", err)
			return err
		}
		return nil
	}
}
