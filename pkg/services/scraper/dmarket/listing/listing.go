package listing

import (
	"context"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/dmarket/listing/pool"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/subscription"
)

type Service struct {
	ctx        context.Context
	Pool       *pool.DmarketPool
	publisher  *subscription.Publisher[pool.ListingItem]
	mu         sync.Mutex
	started    bool
	numWorkers int
	startDelay time.Duration
	cancelFunc context.CancelFunc
}

type ServiceCfg struct {
	Ctx        context.Context
	Pool       *pool.DmarketPool
	StartDelay time.Duration
	NumWorkers int
}

func New(cfg ServiceCfg) *Service {
	return &Service{
		ctx:        cfg.Ctx,
		Pool:       cfg.Pool,
		numWorkers: cfg.NumWorkers,
		startDelay: cfg.StartDelay,
		started:    false,
		publisher:  subscription.NewPublisher[pool.ListingItem](),
		mu:         sync.Mutex{},
	}
}
