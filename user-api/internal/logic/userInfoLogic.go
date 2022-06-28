package logic

import (
	"context"
	"go-zero-rpctest/user/pb"
	"strconv"

	"go-zero-rpctest/user-api/internal/svc"
	"go-zero-rpctest/user-api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	userResp, err := l.svcCtx.UserRpcClient.GetUser(l.ctx, &pb.IdRequest{
		Id: strconv.Itoa(int(req.UserId)),
	})
	if err != nil {
		logx.Error(err)
	}
	userRespId, _ := strconv.Atoi(userResp.Id)
	return &types.UserInfoResp{
		UserId:   int64(userRespId),
		Username: userResp.Name,
	}, nil

}
