package cli

import (
	"log"
	"net/http"
	"os"

	"github.com/coinbase-samples/pay-sdk-go"
	"github.com/spf13/cobra"
)

var Client *pay.Client

var rootCmd = &cobra.Command{
	Use:   "pay",
	Short: "Coinbase Pay enables fast buying and transfering of blockchain tokens using traditional payment methods",
	Long:  `Pay CLI - built for interacting with the Coinbase Pay REST APIs for fast and flexible onramping user funds to the blockchain.`,
}

func Execute() {
	InitClient()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func InitClient() {
	creds, err := pay.SetCredentials()
	if err != nil {
		log.Fatalf("error reading environmental variable: %s", err)
	}

	Client = pay.NewClient(creds, http.Client{})
}
