package handler

import (
	"net/http"

	"greet/internal/logic"
	"greet/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func helloworldHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewHelloworldLogic(r.Context(), svcCtx)
		resp, err := l.Helloworld()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
