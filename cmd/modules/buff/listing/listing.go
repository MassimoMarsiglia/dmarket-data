package bufflisting

import (
	"context"
	"time"

	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules"
	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/config"
	bufflisting "github.com/MassimoMarsiglia/dmarket-bot/pkg/services/scraper/buff/buff_listing"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type Module struct {
	accDir  string
	mapDir  string
	enabled bool
	delay   time.Duration
	srvc    *bufflisting.Service
	logger  *zap.Logger
	nc      *nats.Conn
}

func New(logger *zap.Logger, cmd *cobra.Command, cfg config.BuffListingCfg) (*Module, error) {
	bufflisting := &Module{
		mapDir:  cfg.MappingDir,
		logger:  logger,
		accDir:  cfg.AccDir,
		enabled: true,
		delay:   cfg.Delay,
		nc:      cfg.Conn,
	}
	return bufflisting, nil
}

func (n *Module) Service() *bufflisting.Service {
	return n.srvc
}

// InitFlags implements [Module].
func InitFlags(cmd *cobra.Command) {
	cmd.Flags().Duration("buff.listing.delay", 100*time.Millisecond, "new bufflisting refresh rate")
}

// Name implements [Module].
func (n *Module) Name() string {
	return "buff.listing"
}

// Run implements [Module].
func (n *Module) Run(ctx context.Context) error {
	if !n.enabled {
		return nil
	}

	bufflistingSvc, err := bufflisting.New(bufflisting.ServiceCfg{
		Conn:    n.nc,
		Delay:   n.delay,
		Context: context.Background(),
		Logger:  n.logger,
		AccDir:  &n.accDir,
		MapDir:  n.mapDir,
	})
	if err != nil {
		return err
	}
	n.srvc = bufflistingSvc
	return nil
}

var _ (modules.Module) = (*Module)(nil)
