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
	"log"
	"time"

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

		token, err := Client.GetSessionToken(ctx, d)
		if err != nil {
			fmt.Printf("failed to get session token: %s\n", err)
			return
		}
		resp, err := ResponseToJson(cmd, token)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(resp)
	},
}

func init() {
	rootCmd.AddCommand(sessionCmd)
	sessionCmd.Flags().StringVar(&address, "address", "", "destination wallet address")
	sessionCmd.Flags().StringArrayVar(&blockchains, "chains", nil, "blockchains to allow")
	sessionCmd.Flags().StringArrayVar(&assets, "assets", nil, "tokens to include in the session")
	sessionCmd.Flags().StringArrayVar(&networks, "networks", nil, "networks to allow")
	sessionCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")

	sessionCmd.MarkFlagRequired("address")

}
