package grpc_proxy_middleware

import (
	"github.com/pkg/errors"
	"go_gateway/dao"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"strings"
)

func GrpcHeaderTransferMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream,
	info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, ok := metadata.FromIncomingContext(ss.Context())
		if !ok {
			return errors.New("miss metadata from context")
		}
		for _, item := range strings.Split(serviceDetail.GRPCRule.HeaderTransfor, ",") {
			items := strings.Split(item, " ")
			if len(items) != 3 {
				continue
			}
			if items[0] == "add" || items[0] == "edit" {
				md.Set(items[1], items[2])
			}
			if items[0] == "del" {
				delete(md, items[1])
			}
		}
		//bug: GrpcFlowCountMiddleware failed with error transport: the stream is done or Writ
		//eHeader was already called

		//if err:=ss.SendHeader(md);err != nil {
		//	return errors.WithMessage(err,"SendHeader")
		//}

		if err := handler(srv, ss); err != nil {
			log.Printf("GrpcHeaderTransferMiddleware failed with error %v\n", err)
			return err
		}
		return nil
	}
}
