package scenebreakdown

import (
	"context"

	"github.com/google/uuid"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const namespace = "2f704144-5538-5d38-99ab-e5f6d44478e8"

type SceneBreakdown struct {
	textAnalyzer _interfaces.SceneTextAnalyzer
	uuid         _interfaces.UUIDGenerator
	repository   _interfaces.SceneAnalysisRepository
}

func New(textAnalyzer _interfaces.SceneTextAnalyzer, uuid _interfaces.UUIDGenerator,
	repository _interfaces.SceneAnalysisRepository,
) *SceneBreakdown {
	return &SceneBreakdown{textAnalyzer: textAnalyzer, uuid: uuid, repository: repository}
}

func (ref *SceneBreakdown) BreakdownScene(ctx context.Context,
	tagCategories entity.TagCategories,
	scene entity.Scene,
) (*entity.SceneBreakdown, error) {
	sceneID := generateSceneHash(scene.Text)

	var sceneAnalysis *entity.SceneAnalysis

	sceneAnalysis, err := ref.repository.Get(ctx, sceneID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if sceneAnalysis == nil {
		sceneElements, err := ref.textAnalyzer.AnalyzeSceneText(ctx, scene.Text)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		sceneAnalysis = &entity.SceneAnalysis{
			SceneID:       sceneID,
			SceneElements: sceneElements,
		}

		err = ref.repository.Save(ctx, *sceneAnalysis)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return &entity.SceneBreakdown{
		Number: scene.Number,
		Tags:   ref.generateTags(ctx, sceneAnalysis.SceneElements, tagCategories),
	}, nil
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

func generateSceneHash(scene string) string {
	return uuid.NewMD5(uuid.MustParse(namespace), []byte(scene)).String()
}
