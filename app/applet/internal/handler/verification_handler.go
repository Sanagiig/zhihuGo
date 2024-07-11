package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"githuc.com/Sanagiig/zhihuGo/app/applet/internal/logic"
	"githuc.com/Sanagiig/zhihuGo/app/applet/internal/svc"
	"githuc.com/Sanagiig/zhihuGo/app/applet/internal/types"
)

func VerificationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VerificationRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewVerificationLogic(r.Context(), svcCtx)
		resp, err := l.Verification(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
