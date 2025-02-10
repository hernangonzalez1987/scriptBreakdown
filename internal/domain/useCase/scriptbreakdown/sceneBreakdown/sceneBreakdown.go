package scenebreakdown

import (
	"context"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type SceneBreakdown struct {
	textAnalyzer _interfaces.SceneTextAnalyzer
	uuid         _interfaces.UUIDGenerator
}

func New(textAnalyzer _interfaces.SceneTextAnalyzer, uuid _interfaces.UUIDGenerator) *SceneBreakdown {
	return &SceneBreakdown{textAnalyzer: textAnalyzer, uuid: uuid}
}

func (ref *SceneBreakdown) BreakdownScene(ctx context.Context,
	tagCategories entity.TagCategories,
	scene entity.Scene,
) (*entity.SceneBreakdown, error) {
	elements, err := ref.textAnalyzer.AnalyzeSceneText(ctx, scene.Text)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &entity.SceneBreakdown{Number: scene.Number, Tags: ref.generateTags(ctx, elements, tagCategories)}, nil
}

func (ref *SceneBreakdown) generateTags(ctx context.Context, elementsByCategory map[string][]string,
	tagCategories entity.TagCategories,
) []entity.Tag {
	tags := []entity.Tag{}
	count := 0

	for ctgName, elements := range elementsByCategory {
		for _, element := range elements {
			ctg := tagCategories.GetByType(valueobjects.GetTagCategoryByName(ctgName))
			if ctg == nil {
				log.Ctx(ctx).Warn().Any("categoryName", ctgName).Msg("tag category from text analyzer unknown")

				continue
			}

			count++

			tags = append(tags, entity.Tag{
				ID:       ref.uuid.New().String(),
				Category: *ctg,
				Label:    element,
				Number:   count,
			})
		}
	}

	return tags
}
