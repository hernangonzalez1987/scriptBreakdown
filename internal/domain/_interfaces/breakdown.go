package _interfaces

import (
	"context"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

type ScriptBreakdownUseCase interface {
	BreakdownScript(ctx context.Context,
		event entity.ScriptBreakdownEvent,
	) (result *entity.ScriptBreakdownResult, err error)
}

type SceneBreakdownUseCase interface {
	BreakdownScene(ctx context.Context,
		tagCategories entity.TagCategories,
		scene entity.Scene,
	) (*entity.SceneBreakdown, error)
}

type ScriptBreakdownRequestUseCase interface {
	RequestScriptBreakdown(ctx context.Context, req entity.ScriptBreakdownRequest) (*entity.ScriptBreakdownResult, error)
	GetResult(ctx context.Context, breakdownID string) (
		*entity.ScriptBreakdownResult, error,
	)
}
