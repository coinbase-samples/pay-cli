package cli

import (
	"encoding/json"
	"fmt"

	"github.com/coinbase-samples/pay-sdk-go"
	"github.com/spf13/cobra"
)

var countryId string
var subdivisionId string
var address string
var blockchains []string
var assets []string
var networks []string
var crypto string
var network string
var amount string
var payment string
var method string
var country string
var subid string

func ResponseToJson(cmd *cobra.Command, response interface{}) (string, error) {
	formatBool, err := CheckFormatFlag(cmd)
	if err != nil {
		return "", err
	}
	resp, err := MarshalJSON(response, formatBool)
	if err != nil {
		return "", err
	}
	return string(resp), nil
}

func CheckFormatFlag(cmd *cobra.Command) (bool, error) {
	formatFlagValue, err := cmd.Flags().GetString("format")
	if err != nil {
		return false, fmt.Errorf("cannot read format flag: %w", err)
	}
	return formatFlagValue == "true", nil
}

func MarshalJSON(data interface{}, format bool) ([]byte, error) {
	if format {
		return json.MarshalIndent(data, "", " ")
	}
	return json.Marshal(data)
}

func ConfigToJson(s []byte) (string, error) {

	config := &pay.ConfigData{}
	json.Unmarshal(s, &config)
	j, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		return "", err

	}
	return string(j), nil

}

var SessionDescription = `Onramp Session command returns a unique string that can be passed into the Pay SDK to initialize the onramp widget. 
This token is associated with the Destination Wallets and App ID header provided in the request.

example:
	- pay session --address bc1qrdvlkt8rqsyj229thqzhm0q39edwdj2k7yps6x --chains "bitcoin"`

var ConfigDescription = `View a list of countries supported by Coinbase Pay and the payment methods available in each country. 

examples:

Obtain the list of countries supported by Coinbase Pay, and the payment methods available in each country.
	- pay buyconfig 

Obtain the available payment options for buying Crypto with CBPay:
	- pay buyptions --country US
	- pay buyoptions --country US --sub NY`

var OnrampDescription = `Onramp generates a URL that is launched by a browser as an alternative to integrating Coinbase Pay through initOnRamp.

example:
	- pay onramp --address 0x --chains eth --assets usdc
`

var QuoteDescription = `Returns a quote based on the asset the user would like to purchase, the network they plan to purchase it on, the dollar amount of the payment, the payment currency, the payment method, and country of the user.

example:
	- pay quote --crypto BTC --amount 100.00 --payment USD --method CARD --country US --network Bitcoin --sub NY `

var TransactionDescription = `Provides clients with a list of user CBPay transactions.

exmaple:
	- pay tx --puid 12

`
