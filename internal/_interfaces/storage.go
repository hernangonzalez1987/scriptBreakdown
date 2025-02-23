package _interfaces

import (
	"context"
	"io"
)

type Storage interface {
	Put(ctx context.Context, key string, reader io.Reader) error
	Get(ctx context.Context, key string) (io.ReadCloser, error)
}
