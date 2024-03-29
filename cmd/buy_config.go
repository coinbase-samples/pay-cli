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

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "buyConfig",
	Short: "Returns a list of supported countries, payment methods, currencies, and crypto by Coinbase Pay.",
	Long:  ConfigDescription,
	Run: func(cmd *cobra.Command, args []string) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		resp, err := Client.BuyConfig(ctx)
		if err != nil {
			fmt.Printf("error obtaining buy config %s", err)
			return
		}
		config, err := ConfigToJson(cmd, resp)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Print(config)

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&format, "format", "f", "false", "Pass true for formatted JSON. Default is false")

}
