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

type ScriptBreakdownRequestUseCase struct {
	storage    _interfaces.Storage
	repository _interfaces.BreakdownRepository
}

func New(storage _interfaces.Storage, repository _interfaces.BreakdownRepository) *ScriptBreakdownRequestUseCase {
	return &ScriptBreakdownRequestUseCase{storage: storage, repository: repository}
}

func (ref *ScriptBreakdownRequestUseCase) RequestScriptBreakdown(ctx context.Context,
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

	err = ref.storage.Put(ctx, breakdownID.String(), tempFile)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &entity.ScriptBreakdownResult{BreakdownID: breakdownID.String()}, nil
}

func (ref *ScriptBreakdownRequestUseCase) GetResult(ctx context.Context, breakdownID string) (
	*entity.ScriptBreakdownResult, error,
) {
	result, err := ref.repository.Get(ctx, breakdownID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if result == nil {
		return nil, errors.New("breakdown not found")
	}

	if result.Status != valueobjects.BreakdownStatusSuccess {
		return result, nil
	}

	result.Content, err = ref.storage.Get(ctx, breakdownID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
