package deepseek

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

const path = "/deepseek"

type deepseekIntegrator struct {
	client *http.Client
	host   string
	_interfaces.SceneTextAnalyzer
}

func New(client *http.Client, host string) _interfaces.SceneTextAnalyzer {
	return &deepseekIntegrator{client: client, host: host}
}

func (ref *deepseekIntegrator) AnalyzeSceneText(ctx context.Context, sceneText string) ([]entity.Tag, error) {

	resp, err := ref.client.Get(ref.host + path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	bodyContent := []byte{}
	_, err = resp.Body.Read(bodyContent)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response := DeepSeekContract{}
	err = json.Unmarshal(bodyContent, &response)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return response.ToDomain(), nil
}
