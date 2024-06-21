package demo

import (
	"net/http"

	"gozero_demo/internal/logic/v1/demo"
	"gozero_demo/internal/svc"
	"gozero_demo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
)

// 列表
func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DemoListReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.CodeMsg{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			})
			return
		}

		l := demo.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
