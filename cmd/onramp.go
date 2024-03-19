/**
 * Copyright 2024-present Coinbase Global, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
