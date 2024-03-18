//to do :
// add flags

package cli

import (
	"context"
	"fmt"
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
			fmt.Printf("error obtaining buy config %s", err)
			return
		}
		config, err := ConfigToJson(cmd, resp)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(config)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")

}
