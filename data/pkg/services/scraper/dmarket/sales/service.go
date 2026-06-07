package sales

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/sales"
	"github.com/gammazero/deque"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var ErrNoClients = errors.New("no client running")
var ErrNoItemsInQueue = errors.New("no items in request queue")

type (
	ServiceCfg struct {
		Conn        *nats.Conn
		SalesEm     *sales.Emitter
		DmarketCfgs []dmarket.DmarketCfg
		Delay       time.Duration
		Context     context.Context
		Logger      *zap.Logger
		AccDir      *string
	}

	Service struct {
		nc            *nats.Conn
		em            *sales.Emitter
		logger        *zap.Logger
		filters       []dmarket.FilterFunc[models.Sale]
		context       context.Context
		clients       *deque.Deque[*dmarket.ClientWithResponses]
		queue         *deque.Deque[*dmarket.AggregatorGetLastSalesParams]
		priorityQueue *deque.Deque[*dmarket.AggregatorGetLastSalesParams]

		mu     sync.Mutex
		ticker *time.Ticker
	}
)
