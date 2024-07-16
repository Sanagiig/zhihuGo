package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"githuc.com/Sanagiig/zhihuGo/app/applet/internal/config"
	"githuc.com/Sanagiig/zhihuGo/app/user/rpc/user"
	"githuc.com/Sanagiig/zhihuGo/pkg/interceptors"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	BizRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	bizRedis, err := redis.NewRedis(c.BizRedis)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userRPC),
		BizRedis: bizRedis,
	}
}
