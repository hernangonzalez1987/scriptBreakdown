package ai

import (
	"context"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_ai_AnalyzeSceneText(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	resp := "```json\n{\n  \"props\": [\n    \"mesa de trabalho\",\n\"computador\"\n],\n " +
		"\"cast\": [\n    \"Empregado\"\n]\n}\n```"

	model := NewFakeLLM([]string{
		resp,
	})

	cache := _mocks.NewMockCache[string](t)
	cache.EXPECT().Get("58fb72cd-ece1-583a-9bfb-8ae219141f0f").Return(nil, false)
	cache.EXPECT().Save("58fb72cd-ece1-583a-9bfb-8ae219141f0f", resp)

	llm := New(model, cache)

	scene := `Uma mesa de trabalho com um computador ligado em uma ampla sala 
	de escritório moderna e bem decorada. \nUm EMPREGADO, 
	de camisa social, trabalha focado na tela.`

	tags, err := llm.AnalyzeSceneText(ctx, scene)

	expected := []entity.Tag{
		{Category: "props", Element: "computador"},
		{Category: "props", Element: "mesa de trabalho"},
		{Category: "cast", Element: "Empregado"},
	}

	require.NoError(t, err)
	assert.ElementsMatch(t, expected, tags)
}

func Test_parseResponse(t *testing.T) {
	t.Parallel()

	resp := "```json\n{\n  \"props\": [\n    \"mesa de trabalho\",\n\"computador\"\n],\n" +
		"\"cast\": [\n    \"Empregado\"\n]\n}\n```"

	parsed, err := parseResponse(resp)

	expected := map[string][]string{
		"props": {"mesa de trabalho", "computador"},
		"cast":  {"Empregado"},
	}

	require.NoError(t, err)
	assert.Equal(t, expected, parsed)
}

func Test_extractTagsFromResponse(t *testing.T) {
	t.Parallel()

	parsed := map[string][]string{
		"props": {"mesa de trabalho", "computador"},
		"cast":  {"Empregado"},
	}

	expected := []entity.Tag{
		{Category: "props", Element: "mesa de trabalho"},
		{Category: "props", Element: "computador"},
		{Category: "cast", Element: "Empregado"},
	}

	tags := extractTagsFromResponse(parsed)

	assert.ElementsMatch(t, expected, tags)
}
