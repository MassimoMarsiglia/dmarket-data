package order_book

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client"
	"github.com/gammazero/deque"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var ErrNoClients = errors.New("no client running")
var ErrNoItemsInQueue = errors.New("no items in request queue")

const NATS_KEY = "dmarket.orderbook"

type (
	ServiceCfg struct {
		Conn        *nats.Conn
		DmarketCfgs []client.DmarketCfg
		Delay       time.Duration
		Context     context.Context
		Logger      *zap.Logger
		AccDir      *string
		ItemDir     string
	}

	Service struct {
		nc      *nats.Conn
		logger  *zap.Logger
		filters []client.FilterFunc[any]
		context context.Context
		clients *deque.Deque[*client.ClientWithResponses]
		queue   *deque.Deque[*client.GetOrderBookParams]

		mu     sync.Mutex
		ticker *time.Ticker
	}
)
