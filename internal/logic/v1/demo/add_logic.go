package demo

import (
	"context"
	"net/http"

	"gozero_demo/internal/dal/model"
	"gozero_demo/internal/svc"
	"gozero_demo/internal/types"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 新增
func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.DemoAddReq) (resp *types.DemoAddRes, err error) {
	demoModel := l.svcCtx.DBQuery.Demo

	gender := gofakeit.Gender()
	var sex uint8 = 3
	if gender == "male" {
		sex = 1
	} else if gender == "female" {
		sex = 2
	}

	err = demoModel.WithContext(l.ctx).Create(&model.Demo{
		Name:        gofakeit.Name(),
		Sex:         sex,
		DateOfBirth: gofakeit.Date().Format("2006-01-02"),
		Avatar:      "",
		Email:       gofakeit.Email(),
		Phone:       gofakeit.Phone(),
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("新增失败: %v", err)
		return nil, errors.New(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return
}
