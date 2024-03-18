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

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Generates a Coinbase Pay session key",
	Long:  SessionDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if address == "" {
			log.Fatal("Please specify an address to send funds to with --address flag")
		}

		d := &pay.DestinationWallet{
			Address:           address,
			Blockchains:       &blockchains,
			Assets:            &assets,
			SupportedNetworks: &networks,
		}

		token, err := sdk.Client.GetSessionToken(ctx, d)
		if err != nil {
			fmt.Printf("failed to get session token: %s\n", err)
			return
		}
		TokenToJson(token)
	},
}

func init() {
	rootCmd.AddCommand(sessionCmd)
	sessionCmd.Flags().StringVar(&address, "address", "", "destination wallet address")
	sessionCmd.Flags().StringArrayVar(&blockchains, "chains", nil, "blockchains to allow")
	sessionCmd.Flags().StringArrayVar(&assets, "assets", nil, "tokens to include in the session")
	sessionCmd.Flags().StringArrayVar(&networks, "networks", nil, "networks to allow")
	sessionCmd.MarkFlagRequired("address")

}
