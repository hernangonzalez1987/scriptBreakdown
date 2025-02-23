package llm

import (
	"context"
	"encoding/json"
	"regexp"

	"github.com/google/uuid"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_interfaces"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/tmc/langchaingo/llms"
)

const (
	promptQuestionPart1 = "Find on the following scene, all the elements of the following categories: "
	promptQuestionPart2 = "The answer should be only a JSON, with the elements grouped on arrays by category. " +
		"Every element should be an exact transcription of part of the scene, with no rephrasing. " +
		"Omit on the response the empty categories." +
		"The scene: "
	nameSpace = "2f704144-5538-5d38-99ab-e5f6d44478e8"
)

type Analyzer struct {
	model llms.Model
	cache _interfaces.Cache[string]
}

func New(model llms.Model, cache _interfaces.Cache[string]) *Analyzer {
	return &Analyzer{model: model, cache: cache}
}

func (ref *Analyzer) AnalyzeSceneText(ctx context.Context, sceneText string) (map[string][]string, error) {
	prompt := generatePrompt(sceneText)

	hash := uuid.NewSHA1(uuid.MustParse(nameSpace), []byte(prompt))

	response, exists := ref.cache.Get(hash.String())
	if !exists {
		msgs := []llms.MessageContent{
			{
				Role: llms.ChatMessageTypeHuman,
				Parts: []llms.ContentPart{
					llms.TextPart(prompt),
				},
			},
		}

		resp, err := ref.model.GenerateContent(ctx, msgs)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		if resp == nil || len(resp.Choices) == 0 {
			return nil, errors.New("empty response from model")
		}

		content := resp.Choices[0].Content

		ref.cache.Save(hash.String(), content)

		response = &content
	}

	parsed, err := parseResponse(*response)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Info().Any("parsed", parsed).Msg("response from LLM")

	return parsed, nil
}

func parseResponse(resp string) (map[string][]string, error) {
	exp := regexp.MustCompile(`(?s)[^{]*({.*})[^}]*$`)

	match := exp.FindStringSubmatch(resp)

	jsonResp := map[string][]string{}

	err := json.Unmarshal([]byte(match[1]), &jsonResp)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return jsonResp, nil
}

func generatePrompt(sceneText string) string {
	categoryPrompt := "("
	tagCategoryNames := valueobjects.GetAllTagCategoryNames()

	for i, name := range tagCategoryNames {
		if i > 0 {
			categoryPrompt += ", "
		}
		categoryPrompt += name
	}
	categoryPrompt += "). "

	return promptQuestionPart1 + categoryPrompt + promptQuestionPart2 + sceneText
}
