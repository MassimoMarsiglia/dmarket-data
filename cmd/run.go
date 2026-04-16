package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/NewListing"
	orderbook "github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/OrderBook"
	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/config"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func startNATSServer() *server.Server {
	opts := &server.Options{
		Host: "127.0.0.1",
		Port: 4222,
	}

	s, err := server.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}

	go s.Start()

	if !s.ReadyForConnections(10 * 1e9) {
		log.Fatal("NATS not ready")
	}

	return s
}

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
		s := startNATSServer()
		defer s.Shutdown()

		nc, err := nats.Connect(nats.DefaultURL)
		if err != nil {
			return err
		}

		if cfg.NewListing.Enabled {
			accDir := cfg.NewListing.AccDir
			delay, err := cmd.Flags().GetDuration("dmarket.new_listing.delay")
			if err != nil {
				return err
			}
			logger.Debug("starting with:", zap.Duration("new_listing poll delay:", delay))

			newListingModule, err := NewListing.New(
				logger,
				cmd,
				config.NewListingCfg{
					AccDir: accDir,
					Delay:  delay,
					Conn:   nc,
				},
			)
			if err != nil {
				return err
			}
			newListingModule.Run(context.Background())
		}

		if cfg.OrderBook.Enabled {
			accDir := cfg.OrderBook.AccDir
			itemDir := cfg.OrderBook.ItemDir
			delay, err := cmd.Flags().GetDuration("dmarket.orderbook.delay")
			if err != nil {
				return err
			}
			logger.Debug("starting with:", zap.Duration("new_listing poll delay:", delay))

			orderBookModule, err := orderbook.New(
				logger,
				cmd,
				config.OrderBookCfg{
					ItemDir: itemDir,
					AccDir:  accDir,
					Delay:   delay,
					Conn:    nc,
				},
			)
			if err != nil {
				return err
			}
			orderBookModule.Run(context.Background())
		}

		fmt.Println("running... press Ctrl+C to exit")

		<-ctx.Done() // wait for shutdown signal

		fmt.Println("gracefully shutting down")
		return nil
	},
}
