package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var optionsCmd = &cobra.Command{
	Use:   "buyoptions",
	Short: "Returns a ist of supported countries, payment methods, currencies and crypto by Coinbase Pay",
	Long:  ConfigDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		buyOptions, err := Client.BuyOptions(ctx, countryId, &subdivisionId)
		if err != nil {
			fmt.Printf("error retrieving Buy Options %s", err)
			return
		}
		resp, err := ResponseToJson(cmd, buyOptions)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(resp)

	},
}

func init() {
	rootCmd.AddCommand(optionsCmd)
	optionsCmd.Flags().StringVarP(&countryId, "country", "c", "", "set the country code")
	optionsCmd.Flags().StringVarP(&subdivisionId, "sub", "s", "", "set the subdivision")
	optionsCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")

	optionsCmd.MarkFlagRequired("country")

}
