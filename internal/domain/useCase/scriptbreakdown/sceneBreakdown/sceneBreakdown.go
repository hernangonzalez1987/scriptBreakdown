package scenebreakdown

import (
	"context"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type SceneBreakdown struct {
	textAnalyzer _interfaces.SceneTextAnalyzer
}

func New(textAnalyzer _interfaces.SceneTextAnalyzer) *SceneBreakdown {
	return &SceneBreakdown{textAnalyzer: textAnalyzer}
}

func (ref *SceneBreakdown) BreakdownScene(ctx context.Context, scenes chan entity.Scene,
	sceneBreakdowns chan entity.SceneBreakdown,
) error {
	log.Ctx(ctx).Info().Msg("listening for scenes")

	for scene := range scenes {
		log.Ctx(ctx).Info().Msgf("about to analyze scene %v", scene.Number)

		tags, err := ref.textAnalyzer.AnalyzeSceneText(ctx, scene.Text)
		if err != nil {
			return errors.WithStack(err)
		}

		log.Ctx(ctx).Info().Msgf("finish analyze scene %v", scene.Number)

		sceneBreakdowns <- entity.SceneBreakdown{Number: scene.Number, Tags: tags}
	}

	return nil
}
