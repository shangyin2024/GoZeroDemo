package demo

import (
	"context"
	"gozero_demo/internal/svc"
	"gozero_demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.DemoListReq) (resp *types.DemoListRes, err error) {
	demoModel := l.svcCtx.DBQuery.Demo
	_, _ = demoModel.WithContext(l.ctx).Where(demoModel.ID.Eq("aaa")).Order(demoModel.Phone.Desc()).Take()
	return
}
