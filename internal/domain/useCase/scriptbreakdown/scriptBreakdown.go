package scriptbreakdown

import (
	"context"
	"os"
	"sync"

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
	tempExtention = ".tmp"
)

type BreakdownUseCase struct {
	validate    *validator.Validate
	parser      _interfaces.ScriptParser
	render      _interfaces.ScriptRender
	sceneTagger _interfaces.SceneBreakdown
}

func New(validate *validator.Validate,
	parser _interfaces.ScriptParser,
	render _interfaces.ScriptRender,
	sceneTagger _interfaces.SceneBreakdown,
) *BreakdownUseCase {
	return &BreakdownUseCase{validate: validate, parser: parser, render: render, sceneTagger: sceneTagger}
}

func (ref *BreakdownUseCase) ScriptBreakdown(ctx context.Context,
	req entity.ScriptBreakdownRequest,
) (result *entity.ScriptBreakdownResult, err error) {
	err = ref.validate.Struct(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	inputFile, err := os.Open(req.ScriptFileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() { err = inputFile.Close() }()

	script, err := ref.parser.ParseScript(ctx, inputFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Ctx(ctx).Info().Msgf("script parsed. Number of scenes: %v", len(script.Scenes))

	scriptBreakdown, err := ref.scriptBreakdown(ctx, *script)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tempInputFile, err := os.Open(req.ScriptFileName)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() { err = tempInputFile.Close() }()

	tempOutputFile, err := os.Create(script.Hash + tempExtention)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer func() { err = tempOutputFile.Close() }()

	err = ref.render.RenderScript(ctx, tempInputFile, tempOutputFile, *scriptBreakdown)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &entity.ScriptBreakdownResult{
		BreakdownTempFileName: script.Hash + tempExtention,
	}, nil
}

func (ref *BreakdownUseCase) scriptBreakdown(ctx context.Context,
	script entity.Script,
) (*entity.ScriptBreakdown, error) {
	mu := &sync.Mutex{}
	scriptBreakdown := entity.ScriptBreakdown{SceneBreakdowns: make([]entity.SceneBreakdown, 0)}
	scenes := make(chan entity.Scene)

	group := errgroup.Group{}

	for range numGoRoutines {
		group.Go(func() error {
			for scene := range scenes {
				breakdown, err := ref.sceneTagger.BreakdownScene(ctx, script.TagCategories, scene)
				if err != nil {
					return errors.WithStack(err)
				}
				mu.Lock()
				scriptBreakdown.SceneBreakdowns = append(scriptBreakdown.SceneBreakdowns, *breakdown)
				mu.Unlock()
			}

			return nil
		})
	}

	for _, scene := range script.Scenes {
		scenes <- scene
	}

	close(scenes)

	err := group.Wait()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &scriptBreakdown, nil
}
