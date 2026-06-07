package newlisting

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/dmarket"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/gammazero/deque"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var ErrNoClients = errors.New("no client running")

const NATS_KEY = "dmarket.new_listing"

type (
	ServiceCfg struct {
		Conn        *nats.Conn
		DmarketCfgs []dmarket.DmarketCfg
		Delay       time.Duration
		Context     context.Context
		Logger      *zap.Logger
		AccDir      *string
	}

	Service struct {
		nc      *nats.Conn
		logger  *zap.Logger
		filters []dmarket.FilterFunc[models.Item]
		context context.Context
		clients *deque.Deque[*dmarket.ClientWithResponses]
		mu      sync.Mutex
		ticker  *time.Ticker
	}
)
