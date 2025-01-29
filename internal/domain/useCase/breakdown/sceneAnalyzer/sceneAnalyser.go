package sceneanalyzer

import (
	"context"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

type sceneAnalyzer struct {
	_interfaces.SceneBreakdownTagger
	textAnalyzer _interfaces.SceneTextAnalyzer
}

func New(textAnalyzer _interfaces.SceneTextAnalyzer) _interfaces.SceneBreakdownTagger {
	return &sceneAnalyzer{textAnalyzer: textAnalyzer}
}

func (ref *sceneAnalyzer) BreakdownScene(ctx context.Context, scenes chan entity.Scene,
	sceneBreakdowns chan entity.SceneBreakdown) error {

	for scene := range scenes {
		tags, err := ref.textAnalyzer.AnalyzeSceneText(ctx, scene.Text)
		if err != nil {
			return errors.WithStack(err)
		}
		sceneBreakdowns <- entity.SceneBreakdown{Number: scene.Number, Tags: tags}
	}

	return nil

}
