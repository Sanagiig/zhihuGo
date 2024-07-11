package svc

import (
	"githuc.com/Sanagiig/zhihuGo/app/applet/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
