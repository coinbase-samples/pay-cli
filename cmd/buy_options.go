package cli

import (
	"context"
	"log"
	"time"

	"github.com/coinbase-samples/pay-cli/sdk"
	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:   "buyoptions",
	Short: "Returns a ist of supported countries, payment methods, currencies and crypto by Coinbase Pay",
	Long:  ConfigDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := sdk.Client.BuyOptions(ctx, cid, &sid)
		if err != nil {
			log.Fatalf("error retrieving Buy Options %s", err)
		}
		OptionsToJSON(resp)

	},
}

func init() {
	rootCmd.AddCommand(optionsCmd)
	optionsCmd.Flags().StringVarP(&cid, "country", "c", "", "set the country code")
	optionsCmd.Flags().StringVarP(&sid, "sub", "s", "", "set the subdivision")
}
