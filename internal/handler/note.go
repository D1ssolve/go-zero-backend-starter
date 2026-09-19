package handler

import (
	"errors"
	"net/http"
	"strconv"

	domain "github.com/D1ssolve/go-zero-backend-starter/internal/domain/note"
	"github.com/D1ssolve/go-zero-backend-starter/internal/logic"
	"github.com/D1ssolve/go-zero-backend-starter/internal/svc"
	"github.com/D1ssolve/go-zero-backend-starter/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func createNoteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request types.CreateNoteRequest
		if err := httpx.Parse(r, &request); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		response, err := logic.NewNoteLogic(r.Context(), svcCtx).Create(&request)
		if errors.Is(err, domain.ErrTextRequired) {
			httpx.WriteJsonCtx(r.Context(), w, http.StatusBadRequest, map[string]any{
				"error": map[string]string{"message": err.Error()},
			})
			return
		}
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.WriteJsonCtx(r.Context(), w, http.StatusCreated, response)
	}
}

func listNotesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
		response, err := logic.NewNoteLogic(r.Context(), svcCtx).List(limit)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		httpx.OkJsonCtx(r.Context(), w, response)
	}
}
