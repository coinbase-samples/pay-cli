package cli

import (
	"context"
	"fmt"
	"time"

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

		destinationWallet := pay.DestinationWallet{
			Address:     address,
			Blockchains: &blockchains,
			Assets:      &assets,
		}
		params := pay.OnRampAppParams{
			DestinationWallets: []pay.DestinationWallet{destinationWallet},
		}
		onrampUrl := pay.GenerateOnRampUrlOptions{
			AppId:           Client.Credentials.AppId,
			Host:            &Client.Host,
			OnRampAppParams: params,
		}

		url, err := Client.GenerateOnRampUrl(ctx, onrampUrl)
		if err != nil {
			fmt.Printf("failed to generate onramp URL: %s\n", err)
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
	onrampCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")

	onrampCmd.MarkFlagRequired("address")
	onrampCmd.MarkFlagRequired("blockchains")

}
