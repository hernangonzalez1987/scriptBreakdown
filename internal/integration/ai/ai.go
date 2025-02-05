package ai

import (
	"context"
	"encoding/json"
	"regexp"

	"github.com/google/uuid"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
	"github.com/tmc/langchaingo/llms"
)

const promptQuestion = "Encontra na seguinte cena, todos os elementos cénicos (props, cast, wardrobe, vehicles, animals, sound-effects, music, special-effects, set-dressing). Retorna a resposta em formato JSON, com um array por cada categoria de elemento. A resposta deve ser só o JSON. A cena é a seguinte:"

var nameSpace = uuid.MustParse("2f704144-5538-5d38-99ab-e5f6d44478e8")

type ai struct {
	model llms.Model
	cache _interfaces.Cache
	_interfaces.SceneTextAnalyzer
}

func New(model llms.Model, cache _interfaces.Cache) _interfaces.SceneTextAnalyzer {
	return &ai{model: model, cache: cache}
}

func (ref *ai) AnalyzeSceneText(ctx context.Context, sceneText string) ([]entity.Tag, error) {

	prompt := promptQuestion + sceneText

	hash := uuid.NewSHA1(nameSpace, []byte(prompt))

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

		response = resp.Choices[0].Content

		ref.cache.Save(hash.String(), response)
	}

	parsed, err := parseResponse(response)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return extractTagsFromResponse(parsed), nil

}

func extractTagsFromResponse(parsed map[string][]string) []entity.Tag {

	tags := []entity.Tag{}

	for category, elements := range parsed {
		for _, element := range elements {
			tags = append(tags, entity.Tag{
				Category: category,
				Element:  element,
			})
		}

	}

	return tags
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
