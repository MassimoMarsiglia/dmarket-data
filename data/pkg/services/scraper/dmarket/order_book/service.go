package order_book

import (
	"context"
	"errors"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/gammazero/deque"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var ErrNoClients = errors.New("no client running")
var ErrNoItemsInQueue = errors.New("no items in request queue")

const NATS_KEY = "dmarket.orderbook"

type (
	ServiceCfg struct {
		Conn         *nats.Conn
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
		logger  *zap.Logger
		filters []dmarket.FilterFunc[models.BuyOrder]
		context context.Context
		clients *deque.Deque[*dmarket.ClientWithResponses]
		queue   *deque.Deque[*dmarket.GetOrderBookParams]

		ticker *time.Ticker
	}
)
