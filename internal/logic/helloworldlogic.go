package logic

import (
	"context"

	"greet/internal/svc"
	"greet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloworldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHelloworldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloworldLogic {
	return &HelloworldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HelloworldLogic) Helloworld() (resp *types.HelloWorld, err error) {
	// todo: add your logic here and delete this line

	return &types.HelloWorld{
		Message: "Hello world" + " ",
	}, nil
}
