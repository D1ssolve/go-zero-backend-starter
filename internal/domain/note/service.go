package note

import (
	"context"
	"errors"
	"strings"
)

var ErrTextRequired = errors.New("text is required")

type Repository interface {
	Create(ctx context.Context, text string) (Note, error)
	List(ctx context.Context, limit int) ([]Note, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Create(ctx context.Context, text string) (Note, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return Note{}, ErrTextRequired
	}
	return s.repository.Create(ctx, text)
}

func (s *Service) List(ctx context.Context, limit int) ([]Note, error) {
	if limit < 1 || limit > 100 {
		limit = 20
	}
	return s.repository.List(ctx, limit)
}
