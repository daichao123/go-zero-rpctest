package logic

import (
	"context"

	"go-zero-rpctest/user-api/internal/svc"
	"go-zero-rpctest/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserMiddlewareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserMiddlewareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserMiddlewareLogic {
	return &UserMiddlewareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserMiddlewareLogic) UserMiddleware(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
