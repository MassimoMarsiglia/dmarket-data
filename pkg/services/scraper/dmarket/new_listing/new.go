package newlisting

import (
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client"
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
	svc := &Service{
		nc:      cfg.Conn,
		clients: clients,
		context: cfg.Context,
		ticker:  time.NewTicker(cfg.Delay),
		filters: []client.FilterFunc[models.Item]{client.PriceIDFilter()},
		logger:  cfg.Logger,
	}

	go svc.init()

	once.Do(func() {
		instantiated = &Service{}
	})
	return svc, nil
}
