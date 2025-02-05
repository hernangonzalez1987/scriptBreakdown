package scriptbreakdown

import (
	"context"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_breakdownUseCase_ProcessFile(t *testing.T) {

	ctx := context.Background()
	parser := _mocks.NewMockScriptParser(t)
	sceneTagger := _mocks.NewMockSceneBreakdownTagger(t)

	breakdown := breakdownUseCase{parser: parser, sceneTagger: sceneTagger}

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
		RunAndReturn(func(ctx context.Context, c1 chan entity.Scene, c2 chan entity.SceneBreakdown) error {
			c2 <- entity.SceneBreakdown{Number: 1, Tags: []entity.Tag{{Element: "some"}}}
			return nil
		}).Return(nil)

	result, err := breakdown.ScriptBreakdown(ctx, req)

	var expected *entity.ScriptBreakdownResult

	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mock.AssertExpectationsForObjects(t, parser, sceneTagger)
}
