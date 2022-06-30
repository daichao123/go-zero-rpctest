package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-rpctest/user-api/internal/config"
	"go-zero-rpctest/user-api/internal/middleware"
	"go-zero-rpctest/user-api/model"
	"go-zero-rpctest/user/user"
	"go-zero-rpctest/wallet/wallet"
)

type ServiceContext struct {
	Config          config.Config
	UserMiddleware  rest.Middleware
	UserModel       model.UsersModel
	UserLevelsModel model.UserLevelsModel
	UserRpcClient   user.User
	WalletRpcClient wallet.Wallet
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		UserMiddleware:  middleware.NewUserMiddleware().Handle,
		UserModel:       model.NewUsersModel(sqlx.NewMysql(c.MysqlDb.DataSource)),
		UserLevelsModel: model.NewUserLevelsModel(sqlx.NewMysql(c.MysqlDb.DataSource)),
		//使用user-api作为客户端 调用user-rpc 服务
		UserRpcClient:   user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		WalletRpcClient: wallet.NewWallet(zrpc.MustNewClient(c.WalletRpcConf)),
	}
}
