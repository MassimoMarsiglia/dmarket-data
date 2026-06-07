package order_book

import (
	"context"
	"errors"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/orderbook"
	"github.com/gammazero/deque"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var ErrNoClients = errors.New("no client running")
var ErrNoItemsInQueue = errors.New("no items in request queue")

type (
	ServiceCfg struct {
		Conn         *nats.Conn
		OrderbookEm  *orderbook.Emitter
		DmarketCfgs  []dmarket.DmarketCfg
		Delay        time.Duration
		Context      context.Context
		Logger       *zap.Logger
		AccDir       *string
		SkinsPath    string
		StickersPath string
	}

	Service struct {
		nc      *nats.Conn
		em      *orderbook.Emitter
		logger  *zap.Logger
		filters []dmarket.FilterFunc[models.BuyOrder]
		context context.Context
		clients *deque.Deque[*dmarket.ClientWithResponses]
		queue   *deque.Deque[*dmarket.GetOrderBookParams]

		ticker *time.Ticker
	}
)
