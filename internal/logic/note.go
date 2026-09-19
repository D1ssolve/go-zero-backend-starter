package logic

import (
	"context"
	"time"

	domain "github.com/D1ssolve/go-zero-backend-starter/internal/domain/note"
	"github.com/D1ssolve/go-zero-backend-starter/internal/svc"
	"github.com/D1ssolve/go-zero-backend-starter/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type NoteLogic struct {
	ctx    context.Context
	logger logx.Logger
	svcCtx *svc.ServiceContext
}

func NewNoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoteLogic {
	return &NoteLogic{ctx: ctx, logger: logx.WithContext(ctx), svcCtx: svcCtx}
}

func (l *NoteLogic) Create(request *types.CreateNoteRequest) (*types.Note, error) {
	created, err := l.svcCtx.Notes.Create(l.ctx, request.Text)
	if err != nil {
		return nil, err
	}
	return mapNote(created), nil
}

func (l *NoteLogic) List(limit int) (*types.ListNotesResponse, error) {
	items, err := l.svcCtx.Notes.List(l.ctx, limit)
	if err != nil {
		return nil, err
	}
	result := make([]types.Note, 0, len(items))
	for _, item := range items {
		result = append(result, *mapNote(item))
	}
	return &types.ListNotesResponse{Items: result}, nil
}

func mapNote(value domain.Note) *types.Note {
	return &types.Note{
		ID:        value.ID,
		Text:      value.Text,
		CreatedAt: value.CreatedAt.UTC().Format(time.RFC3339Nano),
	}
}
