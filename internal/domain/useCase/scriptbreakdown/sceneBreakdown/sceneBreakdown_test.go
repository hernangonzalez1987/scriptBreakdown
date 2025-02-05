package scenebreakdown_test

import (
	"context"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	scenebreakdown "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/useCase/scriptbreakdown/sceneBreakdown"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_sceneAnalyzer_AnalyzeScenes(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	scenes := make(chan entity.Scene, 1)
	breakdowns := make(chan entity.SceneBreakdown, 1)

	scenes <- entity.Scene{
		Number: 1,
		Text:   "Some scene",
	}

	aiAnalyzer := _mocks.NewMockSceneTextAnalyzer(t)

	tag := entity.Tag{Element: "Some", Category: "someCategory"}

	aiAnalyzer.EXPECT().AnalyzeSceneText(ctx, "Some scene").Return(
		[]entity.Tag{tag}, nil)

	analyzer := scenebreakdown.New(aiAnalyzer)

	close(scenes)

	err := analyzer.BreakdownScene(context.Background(), scenes, breakdowns)

	breakdown := <-breakdowns

	close(breakdowns)

	require.NoError(t, err)
	assert.Equal(t, entity.SceneBreakdown{Number: 1, Tags: []entity.Tag{tag}}, breakdown)
}
