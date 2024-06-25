package demo

import (
	"context"
	"gozero_demo/internal/common/gormx"
	"gozero_demo/internal/svc"
	"gozero_demo/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
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
	demoExtModel := l.svcCtx.DBQuery.DemoExt

	demos, err := demoModel.WithContext(l.ctx).
		LeftJoin(demoExtModel, demoExtModel.DemoID.EqCol(demoModel.ID), demoExtModel.DeletedAt.IsNull()).
		Scopes(gormx.Paginate(req.Page, req.PageSize)).
		Order(demoModel.Phone.Desc()).
		Find()
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询失败: %v", err)
		return nil, errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	l.svcCtx.RedisClient.Get(l.ctx, "test")

	items := make([]*types.DemoListItem, 0)
	for _, demo := range demos {
		items = append(items, &types.DemoListItem{
			ID:          demo.ID,
			Name:        demo.Name,
			Sex:         demo.Sex,
			DateOfBirth: demo.DateOfBirth,
			Avatar:      demo.Avatar,
			Email:       demo.Email,
			Phone:       demo.Phone,
		})
	}
	return &types.DemoListRes{Items: items}, nil
}
