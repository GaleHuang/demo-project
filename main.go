package main

import (
	"github.com/galehuang/demo-project/config"
	"github.com/galehuang/demo-project/routers"
	"github.com/galehuang/demo-project/services/log"
	"github.com/galehuang/framework-base/osutil"
	"google.golang.org/grpc"
	"net"
	"sync"
)

func NewGRPCServer() *grpc.Server {
	// 如果需要安全证书认证 这里配置
	grpcServer := grpc.NewServer()
	// 注册服务
	routers.RegisterRpcServices(grpcServer)
	return grpcServer
}

func RunAndServeGRPCServer(grpcServer *grpc.Server, address string) {
	// 启动服务
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.GLogger.Critical("start tcp listener err=%v", err)
		return
	}
	defer func() {
		_ = listener.Close()
	}()

	log.GLogger.Critical("grpc server listening at %s", address)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.GLogger.Critical("start grpc server err=%v", err)
		return
	}

}

func main() {
	grpcServer := NewGRPCServer()
	go RunAndServeGRPCServer(grpcServer, config.GetConfig().GRPC.Address)
	// 阻塞等待退出信号
	// 搭配启动脚本 放在最后的"exec ./page-config-project"
	sc := osutil.RegisterExistSignal()
	_ = <-sc

	wait := sync.WaitGroup{}
	// 关闭grpc服务
	wait.Add(1)
	go func() {
		defer wait.Done()
		grpcServer.GracefulStop()
	}()
	// 等待关闭完成
	wait.Wait()

}
