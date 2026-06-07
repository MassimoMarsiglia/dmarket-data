package order_book

import (
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	dmarket_utils "github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/dmarket"
)

var instantiated *Service
var once sync.Once

func New(cfg ServiceCfg) (*Service, error) {
	clients, err := dmarket_utils.LoadDmarketAccounts(cfg.Logger, cfg.AccDir, cfg.DmarketCfgs)
	if err != nil {
		return nil, err
	}
	queue, err := dmarket_utils.LoadDmarketOrderBookQueue(cfg.Logger, cfg.SkinsPath, cfg.StickersPath)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		nc:      cfg.Conn,
		em:      cfg.OrderbookEm,
		clients: clients,
		context: cfg.Context,
		ticker:  time.NewTicker(cfg.Delay),
		filters: []dmarket.FilterFunc[models.BuyOrder]{dmarket.OrderDepthFilter()},
		logger:  cfg.Logger,
		queue:   queue,
	}

	go svc.init()

	once.Do(func() {
		instantiated = &Service{}
	})
	return svc, nil
}
