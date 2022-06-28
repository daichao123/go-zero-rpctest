package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-rpctest/user-api/internal/config"
	"go-zero-rpctest/user-api/internal/middleware"
	"go-zero-rpctest/user-api/model"
)

type ServiceContext struct {
	Config          config.Config
	UserMiddleware  rest.Middleware
	UserModel       model.UsersModel
	UserLevelsModel model.UserLevelsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		UserMiddleware: middleware.NewUserMiddleware().Handle,
	}
}
