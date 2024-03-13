package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coinbase-samples/pay-cli/sdk"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "buy",
	Short: "Returns a ist of supported countries, payment methods, currencies and crypto by Coinbase Pay",
	Long:  ConfigDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if config {
			resp, err := sdk.Client.BuyConfig(ctx)
			if err != nil {
				log.Fatalf("error obtaining buy config %s", err)
			}
			fmt.Printf("buy options: %s", string(resp[:]))
		}
		if options {
			resp, err := sdk.Client.BuyOptions(ctx, cid, &sid)
			if err != nil {
				log.Fatalf("error retrieving Buy Options %s", err)
			}
			OptionsToJSON(resp)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().BoolVar(&config, "config", false, "obtain buy config")
	configCmd.Flags().BoolVar(&options, "options", false, "obtain buy options")
	configCmd.Flags().StringVarP(&cid, "country", "c", "", "set the country code")
	configCmd.Flags().StringVarP(&sid, "sub", "s", "", "set the subdivision")
}
