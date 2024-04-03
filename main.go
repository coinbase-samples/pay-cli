package main

import (
	cli "github.com/coinbase-samples/pay-cli/cmd"

	"github.com/coinbase-samples/pay-cli/sdk"
)

func main() {
	sdk.InitClient()
	cli.Execute()
}
