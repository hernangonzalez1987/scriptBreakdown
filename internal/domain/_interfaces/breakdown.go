package _interfaces

import (
	"context"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

type BreakdownPresentation interface {
	ProcessFile(ctx *gin.Context)
}

type ScriptBreakdownUseCase interface {
	ScriptBreakdown(ctx context.Context, breakdownRequest entity.ScriptBreakdownRequest) (*entity.ScriptBreakdownResult, error)
}

type ScriptParser interface {
	ParseScript(ctx context.Context, reader io.Reader) (*entity.Script, error)
}

type ScriptRender interface {
	RenderScript(ctx context.Context, source io.Reader,
		target io.Writer, breakdown entity.ScriptBreakdown,
	) (err error)
}

type SceneBreakdown interface {
	BreakdownScene(ctx context.Context,
		tagCategories entity.TagCategories,
		scene entity.Scene,
	) (*entity.SceneBreakdown, error)
}

type SceneTextAnalyzer interface {
	AnalyzeSceneText(ctx context.Context, sceneText string) (map[string][]string, error)
}
