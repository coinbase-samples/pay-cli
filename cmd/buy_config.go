package cli

import (
	"context"
	"log"
	"time"

	"github.com/coinbase-samples/pay-cli/sdk"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "buyconfig",
	Short: "Returns a list of supported countries, payment methods, currencies, and crypto by Coinbase Pay.",
	Long:  ConfigDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := sdk.Client.BuyConfig(ctx)
		if err != nil {
			log.Printf("error obtaining buy config %s", err)
			return
		}
		ConfigToJson(resp)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
