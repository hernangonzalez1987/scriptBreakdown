package scriptbreakdown_test

import (
	"context"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_breakdownUseCase_ProcessFile(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	parser := _mocks.NewMockScriptParser(t)
	sceneTagger := _mocks.NewMockSceneBreakdownTagger(t)
	tag := entity.Tag{Element: "some", Category: "someCategory"}
	breakdown := scriptbreakdown.New(validator.New(), parser, sceneTagger)

	req := entity.ScriptBreakdownRequest{
		FilePath: "some input path file",
	}

	script := entity.Script{
		Scenes: []entity.Scene{{Number: 1, Text: "some scene text"}},
	}

	parser.EXPECT().ParseScript(ctx, req).Return(&script, nil)
	sceneTagger.EXPECT().BreakdownScene(ctx,
		mock.AnythingOfType("chan entity.Scene"),
		mock.AnythingOfType("chan entity.SceneBreakdown")).
		RunAndReturn(func(_ context.Context, _ chan entity.Scene, c2 chan entity.SceneBreakdown) error {
			c2 <- entity.SceneBreakdown{Number: 1, Tags: []entity.Tag{tag}}

			return nil
		}).Return(nil)

	result, err := breakdown.ScriptBreakdown(ctx, req)

	expected := &entity.ScriptBreakdownResult{
		FilePath: "someOutputFilePath",
	}

	require.NoError(t, err)
	assert.Equal(t, expected, result)
	mock.AssertExpectationsForObjects(t, parser, sceneTagger)
}
