package _interfaces

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type EventHandler interface {
	HandleEvent(ctx context.Context, s3Event events.S3Event) error
}
