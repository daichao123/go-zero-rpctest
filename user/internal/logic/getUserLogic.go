package logic

import (
	"context"
	"strconv"

	"go-zero-rpctest/user/internal/svc"
	"go-zero-rpctest/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *pb.IdRequest) (*pb.UserResponse, error) {
	// todo: add your logic here and delete this line
	m := map[int]string{
		1: "aaaaa",
		2: "bbbbbb",
	}
	inId, _ := strconv.Atoi(in.Id)
	nickname := "test"
	if s, ok := m[inId]; ok {
		nickname = s
	}
	return &pb.UserResponse{
		Id:   in.Id,
		Name: nickname,
	}, nil
}
