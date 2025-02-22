package scriptbreakdownrequest

import (
	"context"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

const namespace = "1486e36b-d00c-4f29-9098-10eb8eab9002"

type ScriptBreakdownRequestUseCase struct {
	storage _interfaces.Storage
}

func New(storage _interfaces.Storage) *ScriptBreakdownRequestUseCase {
	return &ScriptBreakdownRequestUseCase{storage: storage}
}

func (ref *ScriptBreakdownRequestUseCase) RequestScriptBreakdown(ctx context.Context,
	req entity.ScriptBreakdownRequest,
) (result *entity.ScriptBreakdownResult, err error) {
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
