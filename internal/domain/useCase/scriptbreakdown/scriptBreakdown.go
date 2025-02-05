package scriptbreakdown

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

const (
	numGoRoutines = 2
	bufferSize    = 100
)

type BreakdownUseCase struct {
	validate    *validator.Validate
	parser      _interfaces.ScriptParser
	sceneTagger _interfaces.SceneBreakdownTagger
}

func New(validate *validator.Validate, parser _interfaces.ScriptParser,
	sceneTagger _interfaces.SceneBreakdownTagger,
) *BreakdownUseCase {
	return &BreakdownUseCase{validate: validate, parser: parser, sceneTagger: sceneTagger}
}

func (ref *BreakdownUseCase) ScriptBreakdown(ctx context.Context,
	breakdownRequest entity.ScriptBreakdownRequest,
) (*entity.ScriptBreakdownResult, error) {
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

	scenes := make(chan entity.Scene, bufferSize)
	sceneBreakdowns := make(chan entity.SceneBreakdown, bufferSize)

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
		for range len(script.Scenes) {
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

	return &entity.ScriptBreakdownResult{
		FilePath: "someOutputFilePath",
	}, nil
}
