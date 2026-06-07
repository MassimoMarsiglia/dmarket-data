package orderbook

import (
	"context"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules"
	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/config"
	emorderbook "github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/orderbook"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/dmarket/order_book"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Module struct {
	skinsPath    string
	stickersPath string
	accDir       string
	enabled      bool
	delay        time.Duration
	srvc         *order_book.Service
	logger       *zap.Logger
	nc           *nats.Conn
	orderbookEm  *emorderbook.Emitter
}

func New(logger *zap.Logger, cmd *cobra.Command, cfg config.OrderBookCfg) (*Module, error) {
	orderBook := &Module{
		logger:       logger,
		accDir:       cfg.AccDir,
		skinsPath:    cfg.SkinsPath,
		stickersPath: cfg.StickersPath,
		enabled:      true,
		delay:        cfg.Delay,
		nc:           cfg.Conn,
		orderbookEm:  cfg.OrderBook,
	}
	return orderBook, nil
}

func (m *Module) Service() *order_book.Service {
	return m.srvc
}

// InitFlags implements [Module].
func InitFlags(cmd *cobra.Command) {
	cmd.Flags().Duration("dmarket.orderbook.delay", 100*time.Millisecond, "new listing refresh rate")
}

// Name implements [Module].
func (m *Module) Name() string {
	return "dmarket.new-listing"
}

// Run implements [Module].
func (m *Module) Run(ctx context.Context) error {
	if !m.enabled {
		return nil
	}

	orderBookSvc, err := order_book.New(order_book.ServiceCfg{
		Conn:         m.nc,
		OrderbookEm:  m.orderbookEm,
		SkinsPath:    m.skinsPath,
		StickersPath: m.stickersPath,
		Delay:        m.delay,
		Context:      context.Background(),
		Logger:       m.logger,
		AccDir:       &m.accDir,
	})
	if err != nil {
		return err
	}
	m.srvc = orderBookSvc
	return nil
}

var _ (modules.Module) = (*Module)(nil)
