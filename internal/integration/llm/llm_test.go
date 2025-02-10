package llm

import (
	"context"
	"testing"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	cache.EXPECT().Get(mock.AnythingOfType("string")).Return(nil, false)
	cache.EXPECT().Save(mock.AnythingOfType("string"), resp)

	llm := New(model, cache)

	scene := `Uma mesa de trabalho com um computador ligado em uma ampla sala 
	de escritório moderna e bem decorada. \nUm EMPREGADO, 
	de camisa social, trabalha focado na tela.`

	tags, err := llm.AnalyzeSceneText(ctx, scene)

	expected := map[string][]string{
		"props": {"mesa de trabalho", "computador"},
		"cast":  {"Empregado"},
	}

	require.NoError(t, err)
	assert.Equal(t, expected, tags)
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

func Test_generatePrompt(t *testing.T) {
	t.Parallel()

	type args struct {
		sceneText string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should generate prompt",
			args: args{
				sceneText: "some scene text",
			},
			want: "Find on the following scene, all the elements of the following categories: " +
				"(animals, art department, background actors, camera, cast, makeup/hair, mechanical effects," +
				" music, props, security, set dressing, sound, special effects, special equipment, stunts," +
				" vehicles, visual effects, wardrove). The asnwer should be only a JSON, " +
				"with the elements grouped on arrays by category.The scene: some scene text",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := generatePrompt(tt.args.sceneText); got != tt.want {
				t.Errorf("generatePrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}
