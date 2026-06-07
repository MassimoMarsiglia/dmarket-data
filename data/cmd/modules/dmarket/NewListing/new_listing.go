package NewListing

import (
	"context"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules"
	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/config"
	newlistingem "github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/new_listing"
	newlisting "github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/dmarket/new_listing"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Module struct {
	accDir     string
	enabled    bool
	delay      time.Duration
	srvc       *newlisting.Service
	logger     *zap.Logger
	nc         *nats.Conn
	newListing *newlistingem.Emitter
}

func New(logger *zap.Logger, cmd *cobra.Command, cfg config.NewListingCfg) (*Module, error) {
	newListing := &Module{
		logger:     logger,
		accDir:     cfg.AccDir,
		enabled:    true,
		delay:      cfg.Delay,
		nc:         cfg.Conn,
		newListing: cfg.NewListing,
	}
	return newListing, nil
}

func (n *Module) Service() *newlisting.Service {
	return n.srvc
}

// InitFlags implements [Module].
func InitFlags(cmd *cobra.Command) {
	cmd.Flags().Duration("dmarket.new_listing.delay", 100*time.Millisecond, "new listing refresh rate")
}

// Name implements [Module].
func (n *Module) Name() string {
	return "dmarket.new-listing"
}

// Run implements [Module].
func (n *Module) Run(ctx context.Context) error {
	if !n.enabled {
		return nil
	}

	newListingSvc, err := newlisting.New(newlisting.ServiceCfg{
		Conn:         n.nc,
		NewListingEm: n.newListing,
		Delay:        n.delay,
		Context:      context.Background(),
		Logger:       n.logger,
		AccDir:       &n.accDir,
	})
	if err != nil {
		return err
	}
	n.srvc = newListingSvc
	return nil
}

var _ (modules.Module) = (*Module)(nil)
