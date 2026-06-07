package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	bufflisting "github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/buff/listing"
	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/config"
	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/dmarket/NewListing"
	orderbookmod "github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/dmarket/OrderBook"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/buff/listing"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/new_listing"
	emorderbook "github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/orderbook"
	"github.com/MassimoMarsiglia/dmarket-bot/pkg/nats/emitter/dmarket/sales"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "run cs2 trading bot dataservice",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, stop := signal.NotifyContext(
			context.Background(),
			os.Interrupt,
			syscall.SIGTERM,
		)
		defer stop()

		logger, err := zap.NewDevelopment()
		if err != nil {
			return err
		}
		botConfig, err := config.New[config.BotConfig](cmd, logger)

		cfg, err := botConfig.Load()
		if err != nil {
			return err
		}
		nc, err := nats.Connect()
		if err != nil {
			return err
		}

		orderbookEm := emorderbook.New(emorderbook.Config{Conn: nc})
		newListingEm := newlisting.New(newlisting.Config{Conn: nc})
		_ = sales.New(sales.Config{Conn: nc})
		buffListingEm := listing.New(listing.Config{Conn: nc})

		if cfg.Dmarket.NewListing.Enabled {
			accDir := cfg.Dmarket.NewListing.AccDir
			delay := cfg.Dmarket.NewListing.Delay
			if cmd.Flags().Changed("dmarket.new_listing.delay") {
				delay, err = cmd.Flags().GetDuration("dmarket.new_listing.delay")
				if err != nil {
					return err
				}
			}
			logger.Debug("starting with:", zap.Duration("new_listing poll delay:", delay))

			lgr := logger.With(zap.String("source", "dmarket.new-listing"))

			newListingModule, err := NewListing.New(
				lgr,
				cmd,
				config.NewListingCfg{
					AccDir:     accDir,
					Delay:      delay,
					Conn:       nc,
					NewListing: newListingEm,
				},
			)
			if err != nil {
				return err
			}
			newListingModule.Run(context.Background())
		}

		if cfg.Dmarket.OrderBook.Enabled {
			accDir := cfg.Dmarket.OrderBook.AccDir
			skinsPath := cfg.Dmarket.OrderBook.SkinsPath
			stickersPath := cfg.Dmarket.OrderBook.StickersPath
			delay := cfg.Dmarket.OrderBook.Delay
			if cmd.Flags().Changed("dmarket.orderbook.delay") {
				delay, err = cmd.Flags().GetDuration("dmarket.orderbook.delay")
				if err != nil {
					return err
				}
			}
			logger.Debug("starting with:", zap.Duration("order book poll delay:", delay))

			lgr := logger.With(zap.String("source", "dmarket.order-book"))

			orderBookModule, err := orderbookmod.New(
				lgr,
				cmd,
				config.OrderBookCfg{
					AccDir:       accDir,
					SkinsPath:    skinsPath,
					StickersPath: stickersPath,
					Delay:        delay,
					Conn:         nc,
					OrderBook:    orderbookEm,
				},
			)
			if err != nil {
				return err
			}
			orderBookModule.Run(context.Background())
		}

		if cfg.Buff.Listing.Enabled {
			accDir := cfg.Buff.Listing.AccDir
			mapdir := cfg.Buff.Listing.MappingDir
			delay := cfg.Buff.Listing.Delay
			if cmd.Flags().Changed("buff.listing.delay") {
				delay, err = cmd.Flags().GetDuration("buff.listing.delay")
				if err != nil {
					return err
				}
			}
			logger.Debug("starting with:", zap.Duration("buff listing poll delay:", delay))
			lgr := logger.With(zap.String("source", "buff.listing"))

			buffListingModule, err := bufflisting.New(
				lgr,
				cmd,
				config.BuffListingCfg{
					MappingDir: mapdir,
					AccDir:     accDir,
					Delay:      delay,
					Conn:       nc,
					Listing:    buffListingEm,
				},
			)
			if err != nil {
				return err
			}
			buffListingModule.Run(context.Background())
		}

		// if cfg.Dmarket.Sales.Enabled {
		// 	accDir := cfg.Dmarket.Sales.AccDir
		// 	itemDir := cfg.Dmarket.Sales.ItemDir
		// 	delay, err := cmd.Flags().GetDuration("dmarket.sales.delay")
		// 	if err != nil {
		// 		return err
		// 	}
		// 	logger.Debug("starting with:", zap.Duration("dmarket sales poll delay:", delay))
		// 	lgr := logger.With(zap.String("source", "dmarket.sales"))

		// 	// salesModule, err := sales.New(
		// 	// 	lgr,
		// 	// 	cmd,
		// 	// 	config.SalesCfg{
		// 	// 		ItemDir: itemDir,
		// 	// 		AccDir:  accDir,
		// 	// 		Delay:   delay,
		// 	// 		Conn:    nc,
		// 	// 	},
		// 	// )
		// 	// if err != nil {
		// 	// 	return err
		// 	// }
		// 	// salesModule.Run(context.Background())
		// }

		fmt.Println("running... press Ctrl+C to exit")

		<-ctx.Done() // wait for shutdown signal

		fmt.Println("gracefully shutting down")
		return nil
	},
}
