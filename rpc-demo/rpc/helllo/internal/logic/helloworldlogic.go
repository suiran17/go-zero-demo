package logic

import (
	"context"

	"helllo/helllo"
	"helllo/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type HelloWorldLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHelloWorldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HelloWorldLogic {
	return &HelloWorldLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HelloWorldLogic) HelloWorld(in *helllo.HelloWorldRequest) (*helllo.HelloWorldResponse, error) {
	// todo: add your logic here and delete this line

	return &helllo.HelloWorldResponse{
		Message: "Hello, " + in.Name + "!",
	}, nil
}