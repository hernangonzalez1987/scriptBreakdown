package scriptbreakdown

import (
	"bytes"
	"context"
	"io"
	"sync"
	"time"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

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
) (result *entity.ScriptBreakdownResult, err error) {
	// TODO: Update processing
	// TODO: On error, update on error

	scriptFile, err := ref.sourceStorage.Get(ctx, event.BreakdownID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer func() { err = scriptFile.Close() }()

	current, err := ref.repository.Get(ctx, event.BreakdownID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	version := 1

	if current != nil {
		if current.Status == valueobjects.BreakdownStatusProcessing {
			return nil, errors.New("script is already being processed")
		}
		version = current.Version + 1
	}

	err = ref.repository.Save(ctx, entity.ScriptBreakdownResult{
		BreakdownID:       event.BreakdownID,
		Status:            valueobjects.BreakdownStatusProcessing,
		Version:           version,
		LastUpdate:        time.Now(),
		StatusDescription: "Processing",
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	breakdownResult := entity.ScriptBreakdownResult{
		BreakdownID:       event.BreakdownID,
		Status:            valueobjects.BreakdownStatusProcessing,
		Version:           version + 1,
		LastUpdate:        time.Now(),
		StatusDescription: err.Error(),
	}

	defer func() {
		if err != nil {
			err = ref.repository.Save(ctx, breakdownResult)
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

	breakdownContent := new(bytes.Buffer)
	err = ref.render.RenderScript(ctx, scriptBuffer, breakdownContent, *scriptBreakdown)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	breakdownBuffer := new(bytes.Buffer)
	err = ref.targetStorage.Put(ctx, event.BreakdownID, io.TeeReader(breakdownContent, breakdownBuffer))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	breakdownResult = entity.ScriptBreakdownResult{
		BreakdownID:       event.BreakdownID,
		Status:            valueobjects.BreakdownStatusProcessing,
		Version:           version + 1,
		LastUpdate:        time.Now(),
		StatusDescription: "Success",
	}

	err = ref.repository.Save(ctx, breakdownResult)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &breakdownResult, nil
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
