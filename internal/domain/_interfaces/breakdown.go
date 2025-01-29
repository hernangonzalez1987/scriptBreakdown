package _interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

type BreakdownPresentation interface {
	ProcessFile(ctx *gin.Context)
}

type ScriptBreakdownUseCase interface {
	ProcessFile(ctx context.Context, breakdownRequest entity.ScriptBreakdownRequest) (*entity.ScriptBreakdownResult, error)
}

type ScriptParser interface {
	ParseScript(ctx context.Context, breakdownRequest entity.ScriptBreakdownRequest) (*entity.Script, error)
}

type SceneBreakdownTagger interface {
	BreakdownScene(ctx context.Context, scenes chan entity.Scene, sceneBreakdowns chan entity.SceneBreakdown) error
}

type SceneTextAnalyzer interface {
	AnalyzeSceneText(ctx context.Context, sceneText string) ([]entity.Tag, error)
}
