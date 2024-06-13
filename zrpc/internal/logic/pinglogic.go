package logic

import (
	"context"

	"go-zero-grpc-gateway/common/pb"
	"go-zero-grpc-gateway/zrpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *__.Request) (*__.Response, error) {
	// todo: add your logic here and delete this line
	return &__.Response{Msg: "Pong"}, nil
}
