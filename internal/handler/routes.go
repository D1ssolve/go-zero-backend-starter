package handler

import (
	"net/http"

	"github.com/D1ssolve/go-zero-backend-starter/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, svcCtx *svc.ServiceContext) {
	server.AddRoutes([]rest.Route{
		{Method: http.MethodGet, Path: "/healthz", Handler: healthHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/readyz", Handler: readyHandler(svcCtx)},
		{Method: http.MethodPost, Path: "/api/v1/notes", Handler: createNoteHandler(svcCtx)},
		{Method: http.MethodGet, Path: "/api/v1/notes", Handler: listNotesHandler(svcCtx)},
	})
}
