package logger

import (
	"context"
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Logger struct {
	logger *zerolog.Logger
}

func New() *Logger {
	logger := zerolog.New(os.Stdout).Level(zerolog.InfoLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	return &Logger{logger: &logger}
}

func (ref *Logger) AssociateWithContext(ctx context.Context) context.Context {
	return ref.logger.WithContext(ctx)
}

func (ref *Logger) Logger() *zerolog.Logger {
	return ref.logger
}

func AddCorrelationID(ctx context.Context) context.Context {
	correlationID := uuid.New().String()

	newCtx := zerolog.Ctx(ctx).WithContext(ctx)

	zerolog.Ctx(newCtx).UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("correlation_id", correlationID)
	})

	return newCtx
}
