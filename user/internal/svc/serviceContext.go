package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-rpctest/user-api/model"
	"go-zero-rpctest/user/internal/config"
)

type ServiceContext struct {
	Config          config.Config
	UsersModel      model.UsersModel
	UserLevelsModel model.UserLevelsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UsersModel:      model.NewUsersModel(sqlx.NewMysql(c.MysqlDb.DataSource)),
		UserLevelsModel: model.NewUserLevelsModel(sqlx.NewMysql(c.MysqlDb.DataSource)),
	}
}
