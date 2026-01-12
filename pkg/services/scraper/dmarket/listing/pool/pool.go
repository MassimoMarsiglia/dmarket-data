package pool

import (
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/dmarket/listing/pool/worker"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/pool"
)

type ListingItem struct {
	Name      string
	ItemID    string
	OfferId   *string
	Price     string
	Float     *float32
	PaintSeed *int
	Tradable  bool
}

type DmarketPool struct {
	Pool *pool.Pool[*client.EntityGetItemsResponse, ListingItem]
}

type DmarketPoolConfig struct {
	ReqDelay     time.Duration
	Transformers []pool.TransformFunc
}

func New(cfg DmarketPoolConfig) *DmarketPool {
	transformerFuncs := make([]pool.TransformFunc, 0, len(cfg.Transformers))
	for _, tf := range cfg.Transformers {
		transformerFuncs = append(transformerFuncs, tf)
	}
	return &DmarketPool{
		Pool: pool.New[*client.EntityGetItemsResponse, ListingItem](
			worker.NewDmarketListingWorker(cfg.ReqDelay),
			transformerFuncs,
		),
	}
}
