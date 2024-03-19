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

		r, err := Client.TransactionStatus(ctx, tx)
		if err != nil {
			log.Fatalf("could not get transaction status")
		}
		resp, err := ResponseToJson(cmd, r)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(resp)
	},
}

func init() {
	rootCmd.AddCommand(txCmd)
	txCmd.Flags().StringVar(&puid, "puid", "", "partner user id")
	txCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")

	txCmd.MarkFlagRequired("puid")

}
