package eventhandler

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/_interfaces"
	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
	"github.com/pkg/errors"
)

type S3EventHandler struct {
	breakdownUseCase _interfaces.ScriptBreakdownUseCase
}

func NewS3EventHandler(breakdownUseCase _interfaces.ScriptBreakdownUseCase) *S3EventHandler {
	return &S3EventHandler{breakdownUseCase: breakdownUseCase}
}

func (h *S3EventHandler) HandleEvent(ctx context.Context, s3Event events.S3Event) error {
	for _, record := range s3Event.Records {
		breakdownEvent := entity.ScriptBreakdownEvent{
			BreakdownID: record.S3.Object.Key,
		}

		_, err := h.breakdownUseCase.BreakdownScript(ctx, entity.ScriptBreakdownEvent{
			BreakdownID: breakdownEvent.BreakdownID,
		})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}
