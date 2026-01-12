package worker

import (
	"context"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/subscription"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

var _ WorkerInterface[any] = (*Worker[any])(nil)

type Config[T any] struct {
	ID     uuid.UUID
	Logger *zap.Logger
	Work   WorkFunc[T]
	Delay  time.Duration
}

func New[T any](cfg Config[T]) *Worker[T] {
	return &Worker[T]{
		ID:        cfg.ID,
		logger:    cfg.Logger,
		work:      cfg.Work,
		delay:     cfg.Delay,
		Publisher: subscription.NewPublisher[T](),
	}
}

type WorkFunc[T any] func(ctx context.Context) (T, error)

type Worker[T any] struct {
	ID        uuid.UUID
	delay     time.Duration
	logger    *zap.Logger
	work      WorkFunc[T]
	Publisher *subscription.Publisher[T]
}
