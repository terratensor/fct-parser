package question

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type Question struct {
	ID           uuid.UUID
	DataID       int
	ParentDataID int
	Position     int
	Username     string
	UserRole     string
	Text         string
	Date         time.Time
}

type Store interface {
	Create(ctx context.Context, question Question) (*uuid.UUID, error)
	//Delete(ctx context.Context, question uuid.UUID) error
}

type Questions struct {
	store Store
}

func NewQuestions(store Store) *Questions {
	return &Questions{
		store: store,
	}
}

func (qs *Questions) Create(ctx context.Context, q Question) (*Question, error) {
	q.ID = uuid.New()

	id, err := qs.store.Create(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("create question error: %w", err)
	}
	q.ID = *id
	return &q, nil
}
