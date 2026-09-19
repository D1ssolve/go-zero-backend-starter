package handler

import (
	"net/http"

	"github.com/D1ssolve/go-zero-backend-starter/internal/logic"
	"github.com/D1ssolve/go-zero-backend-starter/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func healthHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := logic.NewHealthLogic(r.Context(), svcCtx).Health()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, response)
	}
}

func readyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := logic.NewHealthLogic(r.Context(), svcCtx).Ready()
		if err != nil {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusServiceUnavailable, map[string]any{
				"error": map[string]string{"message": "postgres is unavailable"},
			})
			return
		}
		httpx.OkJsonCtx(r.Context(), w, response)
	}
}
