package pool

import (
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool/worker"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/subscription"
	"github.com/google/uuid"
)

func New[T, Z any](spawnFunc SpawnFunc[T], TransformFuncs []TransformFunc) *Pool[T, Z] {
	return &Pool[T, Z]{
		Publisher:      subscription.NewPublisher[Z](),
		SpawnFunc:      spawnFunc,
		TransformFuncs: TransformFuncs,
		workers:        make(map[uuid.UUID]worker.Worker[T]),
	}
}
