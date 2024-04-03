package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pay",
	Short: "Coinbase Pay enables fast buying and transfering of blockchain tokens using traditional payment methods",
	Long:  `Pay CLI - built for interacting with the Coinbase Pay REST APIs for fast and flexible onramping user funds to the blockchain.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
