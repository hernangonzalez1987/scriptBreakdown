package _interfaces

import (
	"context"
	"io"
)

type Storage interface {
	Put(ctx context.Context, fileName string, reader io.Reader) error
	Get(ctx context.Context, fileName string, writer io.Writer) error
}
