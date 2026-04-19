package cmd

import (
	"os"

	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules"
	bufflisting "github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/buff/listing"
	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/config"
	"github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/dmarket/NewListing"
	orderbook "github.com/MassimoMarsiglia/dmarket-bot/cmd/modules/dmarket/OrderBook"
	"github.com/spf13/cobra"
)

type Application struct {
	modules []modules.Module
}

var cfg config.Config[config.BotConfig]

func init() {
	// Enable the built-in `completion` command
	rootCmd.CompletionOptions.DisableDefaultCmd = false
	rootCmd.AddCommand(RunCmd)

	config.InitFlags(RunCmd)
	NewListing.InitFlags(RunCmd)
	orderbook.InitFlags(RunCmd)
	bufflisting.InitFlags(RunCmd)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dmarket-bot",
	Short: "cs2 trading bot dataservice",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
