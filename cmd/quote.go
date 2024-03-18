package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/coinbase-samples/pay-cli/sdk"
	"github.com/coinbase-samples/pay-sdk-go"
	"github.com/spf13/cobra"
)

var quoteCmd = &cobra.Command{
	Use:   "quote",
	Short: "Obtain a quote for a crypto purchase",
	Long:  QuoteDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		payload := &pay.BuyQuotePayload{
			PurchaseCurrency: crypto,
			PurchaseNetwork:  &network,
			PaymentAmount:    amount,
			PaymentCurrency:  payment,
			PaymentMethod:    method,
			Country:          country,
			Subdivision:      &subid,
		}

		buyQuote, err := sdk.Client.BuyQuote(ctx, payload)
		if err != nil {
			fmt.Printf("error fetching buy quote: %s", err)
			return
		}
		resp, err := ResponseToJson(cmd, buyQuote)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(resp)
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)

	quoteCmd.Flags().StringVar(&crypto, "crypto", "", "desired token of purchase")
	quoteCmd.Flags().StringVar(&network, "network", "", "desired token of purchase")
	quoteCmd.Flags().StringVar(&amount, "amount", "", "amount of token denoted in local currency set by --payment")
	quoteCmd.Flags().StringVar(&payment, "payment", "USD", "local currency")
	quoteCmd.Flags().StringVar(&method, "method", "", "method of purchase")
	quoteCmd.Flags().StringVar(&country, "country", "", "country of purchase")
	quoteCmd.Flags().StringVar(&subid, "sub", "", "subdivision of the country")

	quoteCmd.MarkFlagRequired("crypto")
	quoteCmd.MarkFlagRequired("amount")
	quoteCmd.MarkFlagRequired("payment")
	quoteCmd.MarkFlagRequired("method")
	quoteCmd.MarkFlagRequired("country")
}
