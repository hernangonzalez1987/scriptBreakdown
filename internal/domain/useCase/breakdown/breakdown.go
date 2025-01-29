package usecasebreakdown

import (
	"context"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

const numGoRoutines = 2

type breakdownUseCase struct {
	parser      _interfaces.ScriptParser
	sceneTagger _interfaces.SceneBreakdownTagger
}

func New(parser _interfaces.ScriptParser, sceneTagger _interfaces.SceneBreakdownTagger) _interfaces.ScriptBreakdownUseCase {
	return &breakdownUseCase{parser: parser, sceneTagger: sceneTagger}
}

func (ref *breakdownUseCase) ProcessFile(ctx context.Context, breakdownRequest entity.ScriptBreakdownRequest) (*entity.ScriptBreakdownResult, error) {

	script, err := ref.parser.ParseScript(ctx, breakdownRequest)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scriptBreakdown := entity.ScriptBreakdown{}

	scenes := make(chan entity.Scene, 10)
	sceneBreakdowns := make(chan entity.SceneBreakdown, 10)

	group := errgroup.Group{}

	for range numGoRoutines {
		group.Go(func() error {
			return ref.sceneTagger.BreakdownScene(ctx, scenes, sceneBreakdowns)
		})
	}

	group.Go(func() error {
		for sceneBreakdown := range sceneBreakdowns {
			scriptBreakdown.SceneBreakdowns[sceneBreakdown.Number] = sceneBreakdown
		}
		return nil
	})

	for _, scene := range script.Scenes {
		scenes <- scene
	}

	close(scenes)
	close(sceneBreakdowns)

	err = group.Wait()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return nil, nil

}
