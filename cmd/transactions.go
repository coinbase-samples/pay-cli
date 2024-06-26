package cli

import (
	"context"
	"log"
	"time"

	"github.com/coinbase-samples/pay-cli/sdk"
	"github.com/coinbase-samples/pay-sdk-go"
	"github.com/spf13/cobra"
)

var puid string
var ps int
var pk string

var txCmd = &cobra.Command{
	Use:   "tx",
	Short: "Real time status of a Pay transaction.",
	Long:  TransactionDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		tx := &pay.TransactionRequest{
			PartnerUserId: puid,
			PageKey:       &pk,
			PageSize:      &ps,
		}

		r, err := sdk.Client.TransactionStatus(ctx, tx)
		if err != nil {
			log.Fatalf("could not get transaction status")
		}
		TransactionToJson(r)
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
	txCmd.Flags().StringVar(&puid, "puid", "", "partner user id")
	txCmd.MarkFlagRequired("puid")

}
