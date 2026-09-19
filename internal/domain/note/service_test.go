package note

import (
	"context"
	"errors"
	"testing"
)

type stubRepository struct {
	createdText string
}

func (s *stubRepository) Create(_ context.Context, text string) (Note, error) {
	s.createdText = text
	return Note{ID: 1, Text: text}, nil
}

func (s *stubRepository) List(_ context.Context, _ int) ([]Note, error) {
	return nil, nil
}

func TestCreate(t *testing.T) {
	repository := &stubRepository{}
	service := NewService(repository)

	created, err := service.Create(context.Background(), "  hello  ")
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if created.Text != "hello" || repository.createdText != "hello" {
		t.Fatalf("Create() text = %q", created.Text)
	}
}

func TestCreateRejectsEmptyText(t *testing.T) {
	service := NewService(&stubRepository{})
	_, err := service.Create(context.Background(), " ")
	if !errors.Is(err, ErrTextRequired) {
		t.Fatalf("Create() error = %v, want %v", err, ErrTextRequired)
	}
}
