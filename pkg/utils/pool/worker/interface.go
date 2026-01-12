package worker

import (
	"context"
)

type WorkerInterface[T any] interface {
	Work(ctx context.Context)
	Start(ctx context.Context) error
}
