package pool

import (
	"context"
	"io"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool/worker"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/subscription"
	"github.com/google/uuid"
)

type (
	SpawnFunc[T any] func(ctx context.Context) (*worker.Worker[T], error)

	TransformFunc func(r io.Reader, w io.Writer) error

	Transformer interface {
		Transform(r io.Reader, w io.Writer) error
	}

	PoolInterface[T, Z any] interface {
		Spawn(ctx context.Context, numWorkers int) ([]worker.Worker[T], error)
		Start(ctx context.Context, startDelay time.Duration) error
		Transform(T) ([]Z, error)
	}

	Pool[T, Z any] struct {
		SpawnFunc      SpawnFunc[T]
		TransformFuncs []TransformFunc
		Publisher      *subscription.Publisher[Z]
		workers        map[uuid.UUID]worker.Worker[T]
	}
)

var _ PoolInterface[any, any] = (*Pool[any, any])(nil)
