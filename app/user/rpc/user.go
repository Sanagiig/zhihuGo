package main

import (
	"flag"
	"fmt"
	cs "github.com/zeromicro/go-zero/core/service"
	"githuc.com/Sanagiig/zhihuGo/pkg/interceptors"

	"githuc.com/Sanagiig/zhihuGo/app/user/rpc/internal/config"
	"githuc.com/Sanagiig/zhihuGo/app/user/rpc/internal/server"
	"githuc.com/Sanagiig/zhihuGo/app/user/rpc/internal/svc"
	"githuc.com/Sanagiig/zhihuGo/app/user/rpc/service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		service.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == cs.DevMode || c.Mode == cs.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 自定义拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
	defer s.Stop()
}
