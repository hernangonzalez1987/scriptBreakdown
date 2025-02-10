package scriptbreakdown

import (
	"context"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_breakdownUseCase_script_breakdown(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	sceneTagger := _mocks.NewMockSceneBreakdown(t)

	breakdown := BreakdownUseCase{sceneTagger: sceneTagger}

	tag := entity.Tag{
		ID: "SomeID", Number: 1, Label: "some",
		Category: entity.Category{Number: 1, ID: "SomeID", Type: valueobjects.TagCategoryProps},
	}

	script := entity.Script{
		Scenes:        []entity.Scene{{Number: 1, Text: "some scene text"}},
		Hash:          "someHash",
		TagCategories: []entity.Category{},
	}

	sceneBreakdown := entity.SceneBreakdown{
		Number: 1,
		Tags:   []entity.Tag{tag},
	}

	sceneTagger.EXPECT().BreakdownScene(ctx, script.TagCategories, script.Scenes[0]).
		Return(&sceneBreakdown, nil)

	result, err := breakdown.scriptBreakdown(ctx, script)

	expected := &entity.ScriptBreakdown{SceneBreakdowns: []entity.SceneBreakdown{
		{
			Number: 1,
			Tags: []entity.Tag{
				{
					ID:     "SomeID",
					Number: 1,
					Category: entity.Category{
						ID:     "SomeID",
						Type:   5,
						Number: 1,
					},
					Label: "some",
				},
			},
		},
	}}

	require.NoError(t, err)
	assert.Equal(t, expected, result)
	mock.AssertExpectationsForObjects(t, sceneTagger)
}
