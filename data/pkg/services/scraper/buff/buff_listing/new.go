package bufflisting

import (
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/buff"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	buff_utils "github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/buff"
)

var instantiated *Service
var once sync.Once

func New(cfg ServiceCfg) (*Service, error) {
	clients, err := buff_utils.LoadBuffAccounts(cfg.Logger, cfg.AccDir, cfg.BuffCfgs)
	if err != nil {
		return nil, err
	}

	queue, err := buff_utils.LoadBuffQueue(cfg.Logger, cfg.MapDir)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		queue:   queue,
		nc:      cfg.Conn,
		clients: clients,
		context: cfg.Context,
		ticker:  time.NewTicker(cfg.Delay),
		filters: []buff.FilterFunc[models.Item]{buff.PriceIDFilter()},
		logger:  cfg.Logger,
	}

	go svc.init()

	once.Do(func() {
		instantiated = &Service{}
	})
	return nil, nil
}
