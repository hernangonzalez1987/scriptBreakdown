package scriptbreakdownrequest

import (
	"context"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	valueobjects "github.com/hernangonzalez1987/scriptBreakdown/internal/domain/valueObjects"
	"github.com/pkg/errors"
)

const namespace = "1486e36b-d00c-4f29-9098-10eb8eab9002"

var errNotFound = valueobjects.NewCustomError("NOTFOUND", "breakdown not found")

type UseCase struct {
	scriptsStorage    _interfaces.Storage
	breakdownsStorage _interfaces.Storage
	repository        _interfaces.BreakdownRepository
}

func New(scriptsStorage _interfaces.Storage,
	breakdownsStorage _interfaces.Storage,
	repository _interfaces.BreakdownRepository,
) *UseCase {
	return &UseCase{scriptsStorage: scriptsStorage, breakdownsStorage: breakdownsStorage, repository: repository}
}

func (ref *UseCase) RequestScriptBreakdown(ctx context.Context,
	req entity.ScriptBreakdownRequest,
) (*entity.ScriptBreakdownResult, error) {
	tempFile, err := os.CreateTemp("", "tempFile")
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer tempFile.Close()

	scriptRawContent, err := io.ReadAll(io.TeeReader(req.TempScriptFile, tempFile))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	breakdownID := uuid.NewMD5(uuid.MustParse(namespace), scriptRawContent)

	err = tempFile.Sync()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	_, err = tempFile.Seek(0, 0)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = ref.scriptsStorage.Put(ctx, breakdownID.String(), tempFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &entity.ScriptBreakdownResult{BreakdownID: breakdownID.String()}, nil
}

func (ref *UseCase) GetResult(ctx context.Context, breakdownID string) (
	*entity.ScriptBreakdownResult, error,
) {
	result, err := ref.repository.Get(ctx, breakdownID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result == nil {
		return nil, errNotFound
	}

	if result.Status != valueobjects.BreakdownStatusSuccess {
		return result, nil
	}

	result.Content, err = ref.breakdownsStorage.Get(ctx, breakdownID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
