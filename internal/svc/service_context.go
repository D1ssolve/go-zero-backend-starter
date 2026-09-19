package svc

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/D1ssolve/go-zero-backend-starter/internal/config"
	"github.com/D1ssolve/go-zero-backend-starter/internal/domain/note"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ServiceContext struct {
	Config config.Config
	DB     *pgxpool.Pool
	Notes  *note.Service
}

func NewServiceContext(ctx context.Context, c config.Config) (*ServiceContext, error) {
	if strings.TrimSpace(c.Postgres.DSN) == "" {
		return nil, errors.New("Postgres.DSN or DATABASE_URL is required")
	}

	pool, err := pgxpool.New(ctx, c.Postgres.DSN)
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, err
	}

	repository := note.NewPostgresRepository(pool)
	return &ServiceContext{
		Config: c,
		DB:     pool,
		Notes:  note.NewService(repository),
	}, nil
}

func (s *ServiceContext) Close() {
	if s.DB != nil {
		s.DB.Close()
	}
}
