package routers

import (
	"github.com/galehuang/demo-project/api/business"
	"github.com/galehuang/demo-project/rpcs"
	"google.golang.org/grpc"
)

func RegisterRpcServices(server *grpc.Server) {
	business.RegisterProductServiceServer(server, &rpcs.ProductServiceServerImpl{})
}
