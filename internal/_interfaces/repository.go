package _interfaces

import (
	"context"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

type BreakdownRepository interface {
	Init(ctx context.Context) error
	Save(ctx context.Context, result entity.ScriptBreakdownResult) error
	Get(ctx context.Context, id string) (*entity.ScriptBreakdownResult, error)
}
