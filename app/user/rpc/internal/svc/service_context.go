package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"githuc.com/Sanagiig/zhihuGo/app/user/rpc/internal/config"
	"githuc.com/Sanagiig/zhihuGo/app/user/rpc/model"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
