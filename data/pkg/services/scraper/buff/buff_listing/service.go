package bufflisting

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/pkg/client/buff"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/models"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/buff/listing"
	buff_utils "github.com/MassimoMarsiglia/dmarket-bot/pkg/utils/buff"
	"github.com/gammazero/deque"
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

var ErrNoClients = errors.New("no client running")
var ErrNoItemsInQueue = errors.New("no items in request queue")

type (
	ServiceCfg struct {
		Conn      *nats.Conn
		ListingEm *listing.Emitter
		BuffCfgs  []buff.BuffCfg
		Delay     time.Duration
		Context   context.Context
		Logger    *zap.Logger
		AccDir    *string
		MapDir    string
	}

	Service struct {
		nc      *nats.Conn
		em      *listing.Emitter
		logger  *zap.Logger
		filters []buff.FilterFunc[models.Item]
		context context.Context
		clients *deque.Deque[*buff.ClientWithResponses]
		queue   *deque.Deque[*buff_utils.GetListingParams]
		mu      sync.Mutex
		ticker  *time.Ticker
	}
)
