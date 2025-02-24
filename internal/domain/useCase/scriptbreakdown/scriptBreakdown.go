package scriptbreakdown

import (
	"bytes"
	"context"
	"io"
	"os"
	"sync"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

var errAlreadyProcessing = errors.New("script is already being processed")

const (
	numGoRoutines = 2
	bufferSize    = 100
)

type BreakdownUseCase struct {
	parser        _interfaces.ScriptParser
	render        _interfaces.ScriptRender
	sceneTagger   _interfaces.SceneBreakdownUseCase
	sourceStorage _interfaces.Storage
	targetStorage _interfaces.Storage
	repository    _interfaces.BreakdownRepository
}

func New(parser _interfaces.ScriptParser,
	render _interfaces.ScriptRender,
	sceneTagger _interfaces.SceneBreakdownUseCase,
	sourceStorage _interfaces.Storage,
	targetStorage _interfaces.Storage,
	repository _interfaces.BreakdownRepository,
) *BreakdownUseCase {
	return &BreakdownUseCase{
		parser:        parser,
		render:        render,
		sceneTagger:   sceneTagger,
		sourceStorage: sourceStorage,
		targetStorage: targetStorage,
		repository:    repository,
	}
}

func (ref *BreakdownUseCase) BreakdownScript(ctx context.Context,
	event entity.ScriptBreakdownEvent,
) (*entity.ScriptBreakdownResult, error) {
	log.Ctx(ctx).Info().Any("event", event).Msg("processing script breakdown")

	var err error

	scriptFile, err := ref.sourceStorage.Get(ctx, event.BreakdownID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer scriptFile.Close()

	current, err := ref.repository.Get(ctx, event.BreakdownID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	version := 1

	if current != nil {
		if current.Status == valueobjects.BreakdownStatusProcessing {
			return nil, errAlreadyProcessing
		}

		if current.Status == valueobjects.BreakdownStatusSuccess {
			log.Ctx(ctx).Info().Any("event", event).Msg("script already processed")

			return nil, nil
		}

		version = current.Version + 1
	}

	err = ref.repository.Save(ctx, entity.ScriptBreakdownResult{
		BreakdownID:       event.BreakdownID,
		Status:            valueobjects.BreakdownStatusProcessing,
		Version:           version,
		StatusDescription: "Processing",
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() {
		if err != nil {
			breakdownResult := entity.ScriptBreakdownResult{
				BreakdownID:       event.BreakdownID,
				Status:            valueobjects.BreakdownStatusError,
				Version:           version + 1,
				StatusDescription: err.Error(),
			}

			err := ref.repository.Save(ctx, breakdownResult)
			if err != nil {
				log.Ctx(ctx).Err(err).Msg("error writing error on db")
			}
		}
	}()

	scriptBuffer := new(bytes.Buffer)

	script, err := ref.parser.ParseScript(ctx, io.TeeReader(scriptFile, scriptBuffer))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	scriptBreakdown, err := ref.scriptBreakdown(ctx, *script)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	log.Ctx(ctx).Info().Any("event", event).Msg("script breakdown done. About to render")

	breakdownContent := new(bytes.Buffer)

	err = ref.render.RenderScript(ctx, scriptBuffer, breakdownContent, *scriptBreakdown)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	tempFile, err := os.CreateTemp("", "tempFile")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, breakdownContent)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = tempFile.Sync()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = tempFile.Seek(0, 0)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = ref.targetStorage.Put(ctx, event.BreakdownID, tempFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	breakdownResult := entity.ScriptBreakdownResult{
		BreakdownID:       event.BreakdownID,
		Status:            valueobjects.BreakdownStatusSuccess,
		Version:           version + 1,
		StatusDescription: "Success",
	}

	log.Ctx(ctx).Info().Any("event", event).Msg("render done")

	err = ref.repository.Save(ctx, breakdownResult)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &breakdownResult, nil
}

func (ref *BreakdownUseCase) scriptBreakdown(ctx context.Context,
	script entity.Script,
) (*entity.ScriptBreakdown, error) {
	mutex := &sync.Mutex{}

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

				mutex.Lock()
				scriptBreakdown.SceneBreakdowns = append(scriptBreakdown.SceneBreakdowns, *breakdown)
				mutex.Unlock()
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
