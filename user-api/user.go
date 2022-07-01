package main

import (
	"flag"
	"fmt"
	"go-zero-rpctest/user-api/internal/config"
	"go-zero-rpctest/user-api/internal/handler"
	"go-zero-rpctest/user-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	//httpx.SetErrorHandler(func(err error) (int, interface{}) {
	//	switch e := err.(type) {
	//	case *xerr.CodeError:
	//		return int(e.GetErrCode()), nil
	//	default:
	//		return http.StatusInternalServerError, err.Error()
	//	}
	//})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
