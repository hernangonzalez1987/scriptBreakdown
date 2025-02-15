package _interfaces

import (
	"context"
	"io"

	"github.com/hernangonzalez1987/scriptBreakdown/internal/domain/entity"
)

type ScriptParser interface {
	ParseScript(ctx context.Context, reader io.Reader) (*entity.Script, error)
}

type ScriptRender interface {
	RenderScript(ctx context.Context, source io.Reader,
		target io.Writer, breakdown entity.ScriptBreakdown,
	) (err error)
}
