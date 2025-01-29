package sceneanalyzer

import (
	"context"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func Test_sceneAnalyzer_AnalyzeScenes(t *testing.T) {

	ctx := context.Background()

	scenes := make(chan entity.Scene, 1)
	breakdowns := make(chan entity.SceneBreakdown, 1)

	scenes <- entity.Scene{
		Number: 1,
		Text:   "Some scene",
	}

	aiAnalyzer := _mocks.NewMockSceneTextAnalyzer(t)

	aiAnalyzer.EXPECT().AnalyzeSceneText(ctx, "Some scene").Return(
		[]entity.Tag{{Name: "Some"}}, nil)

	analyzer := sceneAnalyzer{textAnalyzer: aiAnalyzer}

	close(scenes)

	err := analyzer.BreakdownScene(context.Background(), scenes, breakdowns)

	breakdown := <-breakdowns

	close(breakdowns)

	assert.NoError(t, err)
	assert.Equal(t, entity.SceneBreakdown{Number: 1, Tags: []entity.Tag{{Name: "Some"}}}, breakdown)

}
