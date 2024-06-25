package demo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	xhttp "github.com/zeromicro/x/http"
	"gozero_demo/internal/logic/v1/demo"
	"gozero_demo/internal/svc"
	"gozero_demo/internal/types"
)

// 删除
func DelHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DemoDelReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.CodeMsg{
				Code: http.StatusBadRequest,
				Msg:  err.Error(),
			})
			return
		}

		l := demo.NewDelLogic(r.Context(), svcCtx)
		resp, err := l.Del(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
