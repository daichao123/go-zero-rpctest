package logic

import (
	"context"

	"go-zero-rpctest/wallet/internal/svc"
	"go-zero-rpctest/wallet/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCoinInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCoinInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCoinInfoLogic {
	return &GetCoinInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCoinInfoLogic) GetCoinInfo(in *pb.CoinIdRequest) (*pb.WalletResponse, error) {
	// todo: add your logic here and delete this line

	return &pb.WalletResponse{}, nil
}
