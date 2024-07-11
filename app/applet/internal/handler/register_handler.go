package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"githuc.com/Sanagiig/zhihuGo/app/applet/internal/logic"
	"githuc.com/Sanagiig/zhihuGo/app/applet/internal/svc"
	"githuc.com/Sanagiig/zhihuGo/app/applet/internal/types"
)

func RegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
