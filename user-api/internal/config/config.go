package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	RpcClientConf zrpc.RpcClientConf
	MysqlDb       struct {
		DataSource string
	}
}
