package scriptbreakdown

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

const numGoRoutines = 2

type breakdownUseCase struct {
	validate    *validator.Validate
	parser      _interfaces.ScriptParser
	sceneTagger _interfaces.SceneBreakdownTagger
}

func New(validate *validator.Validate, parser _interfaces.ScriptParser, sceneTagger _interfaces.SceneBreakdownTagger) _interfaces.ScriptBreakdownUseCase {
	return &breakdownUseCase{validate: validate, parser: parser, sceneTagger: sceneTagger}
}

func (ref *breakdownUseCase) ScriptBreakdown(ctx context.Context, breakdownRequest entity.ScriptBreakdownRequest) (*entity.ScriptBreakdownResult, error) {

	err := ref.validate.Struct(breakdownRequest)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	script, err := ref.parser.ParseScript(ctx, breakdownRequest)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Ctx(ctx).Info().Msgf("script parsed. Number of scenes: %v", len(script.Scenes))

	scriptBreakdown := entity.ScriptBreakdown{SceneBreakdowns: make([]entity.SceneBreakdown, 0)}

	scenes := make(chan entity.Scene, 100)
	sceneBreakdowns := make(chan entity.SceneBreakdown, 100)

	group := errgroup.Group{}

	for range numGoRoutines {
		group.Go(func() error {
			return ref.sceneTagger.BreakdownScene(ctx, scenes, sceneBreakdowns)
		})
	}

	for _, scene := range script.Scenes {
		scenes <- scene
	}

	go func() {
		for i := 0; i < len(script.Scenes); i++ {

			sceneBreakdown := <-sceneBreakdowns
			scriptBreakdown.SceneBreakdowns = append(scriptBreakdown.SceneBreakdowns, sceneBreakdown)
		}
	}()

	close(scenes)

	err = group.Wait()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	close(sceneBreakdowns)

	fmt.Println(scriptBreakdown)

	return nil, nil

}
