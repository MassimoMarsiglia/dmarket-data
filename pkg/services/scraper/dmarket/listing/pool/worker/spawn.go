package worker

import (
	"context"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool/worker"
	"github.com/google/uuid"
)

func NewDmarketListingWorker(reqDelay time.Duration) pool.SpawnFunc[*client.EntityGetItemsResponse] {
	return func(ctx context.Context) (*worker.Worker[*client.EntityGetItemsResponse], error) {
		id := uuid.New()
		w, err := New(DmarketWorkerConfig{
			ID:       id,
			Dmarket:  client.DmarketCfg{},
			ReqDelay: reqDelay,
		})
		if err != nil {
			return nil, err
		}
		return w.w, nil
	}
}
