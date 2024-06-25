// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	v1demo "gozero_demo/internal/handler/v1/demo"
	"gozero_demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.RequestLogMiddleware, serverCtx.ParameterValidationMiddleware},
			[]rest.Route{
				{
					// 新增
					Method:  http.MethodPost,
					Path:    "/add",
					Handler: v1demo.AddHandler(serverCtx),
				},
				{
					// 删除
					Method:  http.MethodDelete,
					Path:    "/del",
					Handler: v1demo.DelHandler(serverCtx),
				},
				{
					// 列表
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: v1demo.ListHandler(serverCtx),
				},
				{
					// 修改
					Method:  http.MethodPut,
					Path:    "/update",
					Handler: v1demo.UpdateHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/v1/demo"),
	)
}
