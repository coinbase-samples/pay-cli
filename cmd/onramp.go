package cli

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coinbase-samples/pay-cli/sdk"
	"github.com/coinbase-samples/pay-sdk-go"
	"github.com/spf13/cobra"
)

var onrampCmd = &cobra.Command{

	Use:   "onramp",
	Short: "Generates an onramp URL",
	Long:  OnrampDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		d := pay.DestinationWallet{
			Address:     address,
			Blockchains: &blockchains,
			Assets:      &assets,
		}
		p := pay.OnRampAppParams{
			DestinationWallets: []pay.DestinationWallet{d},
		}
		o := pay.GenerateOnRampUrlOptions{
			AppId:           sdk.Client.Credentials.AppId,
			Host:            &sdk.Client.Host,
			OnRampAppParams: p,
		}

		url, err := sdk.Client.GenerateOnRampUrl(ctx, o)
		if err != nil {
			log.Print(err)
			return
		}

		fmt.Println("Generated URL:", url)
	},
}

func init() {
	rootCmd.AddCommand(onrampCmd)
	onrampCmd.Flags().StringVar(&address, "address", "", "destination wallet address")
	onrampCmd.Flags().StringArrayVar(&blockchains, "blockchains", nil, "blockchains to allow")
	onrampCmd.Flags().StringArrayVar(&assets, "assets", nil, "blockchains to allow")

	onrampCmd.MarkFlagRequired("address")
	onrampCmd.MarkFlagRequired("blockchains")

}
