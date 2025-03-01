// Code generated by goctl. DO NOT EDIT.
// Source: demo.proto

package server

import (
	"context"

	"gozero/rpc/demo/demo"
	"gozero/rpc/demo/internal/logic"
	"gozero/rpc/demo/internal/svc"
)

type DemoServer struct {
	svcCtx *svc.ServiceContext
	demo.UnimplementedDemoServer
}

func NewDemoServer(svcCtx *svc.ServiceContext) *DemoServer {
	return &DemoServer{
		svcCtx: svcCtx,
	}
}

func (s *DemoServer) Ping(ctx context.Context, in *demo.Request) (*demo.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
