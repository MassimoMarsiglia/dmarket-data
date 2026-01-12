package pool

import (
	"context"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool/worker"
)

func (p *Pool[T, Z]) Spawn(ctx context.Context, numWorkers int) ([]worker.Worker[T], error) {
	// Implementation specific to the pool type
	workers := make([]worker.Worker[T], 0, numWorkers)
	for i := 0; i < numWorkers; i++ {
		worker, err := p.SpawnFunc(ctx)
		if err != nil {
			return nil, err
		}
		p.workers[worker.ID] = *worker
		workers = append(workers, *worker)
	}
	return workers, nil
}
