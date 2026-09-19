package logic

import (
	"context"
	"time"

	"github.com/D1ssolve/go-zero-backend-starter/internal/svc"
	"github.com/D1ssolve/go-zero-backend-starter/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type HealthLogic struct {
	ctx    context.Context
	logger logx.Logger
	svcCtx *svc.ServiceContext
}

func NewHealthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthLogic {
	return &HealthLogic{ctx: ctx, logger: logx.WithContext(ctx), svcCtx: svcCtx}
}

func (l *HealthLogic) Health() (*types.HealthResponse, error) {
	return &types.HealthResponse{Status: "ok"}, nil
}

func (l *HealthLogic) Ready() (*types.HealthResponse, error) {
	ctx, cancel := context.WithTimeout(l.ctx, 2*time.Second)
	defer cancel()
	if err := l.svcCtx.DB.Ping(ctx); err != nil {
		l.logger.Errorf("postgres readiness check failed: %v", err)
		return nil, err
	}
	return &types.HealthResponse{Status: "ready"}, nil
}
