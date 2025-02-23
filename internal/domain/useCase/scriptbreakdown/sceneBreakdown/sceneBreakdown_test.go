package scenebreakdown_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	scenebreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown/sceneBreakdown"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_sceneAnalyzer_AnalyzeScenes_should_analyse_and_update_db(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	scene := entity.Scene{
		Number: 1,
		Text:   "Some scene",
	}

	aiAnalyzer := _mocks.NewMockSceneTextAnalyzer(t)
	sceneAnalysisRepository := _mocks.NewMockSceneAnalysisRepository(t)

	result := map[string][]string{
		"props": {"someElement"},
	}

	aiAnalyzer.EXPECT().AnalyzeSceneText(ctx, "Some scene").Return(
		result, nil)

	uuidGenerator := _mocks.NewMockUUIDGenerator(t)
	uuidGenerator.EXPECT().New().Return(uuid.MustParse("4080f145-f3aa-4d82-82d6-fccbde230e1f"))
	sceneAnalysisRepository.EXPECT().Get(ctx, "ae65ccfd-d30d-33a8-bec0-1832cc071984").Return(nil, nil)
	sceneAnalysisRepository.EXPECT().Save(ctx, entity.SceneAnalysis{
		SceneID:       "ae65ccfd-d30d-33a8-bec0-1832cc071984",
		SceneElements: result,
	}).Return(nil)

	analyzer := scenebreakdown.New(aiAnalyzer, uuidGenerator, sceneAnalysisRepository)

	categories := entity.TagCategories{{
		Type:   valueobjects.TagCategoryProps,
		ID:     "SomeID",
		Number: 1,
	}}

	breakdown, err := analyzer.BreakdownScene(context.Background(), categories, scene)

	expected := &entity.SceneBreakdown{
		Number: 1,
		Tags: []entity.Tag{
			{
				Category: entity.Category{
					ID:     "SomeID",
					Type:   valueobjects.TagCategoryProps,
					Number: 1,
				},
				Label:  "someElement",
				Number: 1,
				ID:     "4080f145-f3aa-4d82-82d6-fccbde230e1f",
			},
		},
	}

	require.NoError(t, err)
	assert.Equal(t, expected, breakdown)
}

func Test_sceneAnalyzer_AnalyzeScenes_should_reuse_existing_analysis(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	scene := entity.Scene{
		Number: 1,
		Text:   "Some scene",
	}

	aiAnalyzer := _mocks.NewMockSceneTextAnalyzer(t)
	sceneAnalysisRepository := _mocks.NewMockSceneAnalysisRepository(t)

	result := map[string][]string{
		"props": {"someElement"},
	}

	uuidGenerator := _mocks.NewMockUUIDGenerator(t)
	uuidGenerator.EXPECT().New().Return(uuid.MustParse("4080f145-f3aa-4d82-82d6-fccbde230e1f"))
	sceneAnalysisRepository.EXPECT().Get(ctx, "ae65ccfd-d30d-33a8-bec0-1832cc071984").
		Return(&entity.SceneAnalysis{
			SceneID:       "ae65ccfd-d30d-33a8-bec0-1832cc071984",
			SceneElements: result,
		}, nil)

	analyzer := scenebreakdown.New(aiAnalyzer, uuidGenerator, sceneAnalysisRepository)

	categories := entity.TagCategories{{
		Type:   valueobjects.TagCategoryProps,
		ID:     "SomeID",
		Number: 1,
	}}

	breakdown, err := analyzer.BreakdownScene(context.Background(), categories, scene)

	expected := &entity.SceneBreakdown{
		Number: 1,
		Tags: []entity.Tag{
			{
				Category: entity.Category{
					ID:     "SomeID",
					Type:   valueobjects.TagCategoryProps,
					Number: 1,
				},
				Label:  "someElement",
				Number: 1,
				ID:     "4080f145-f3aa-4d82-82d6-fccbde230e1f",
			},
		},
	}

	require.NoError(t, err)
	assert.Equal(t, expected, breakdown)
}
