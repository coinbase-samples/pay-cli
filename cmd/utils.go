package cli

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/coinbase-samples/pay-sdk-go"
)

var cid string
var sid string
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

func OptionsToJSON(r *pay.BuyOptionsResponse) {
	j, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		log.Fatalf("error marshalling response into JSON: %s \n. pay response: %v", err, r)
	}
	fmt.Printf("options: %s", string(j))
}

func QuoteToJson(r *pay.BuyQuoteResponse) {
	q, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		log.Fatalf("error marshalling response into JSON: %s \n. pay response: %v", err, r)
	}
	fmt.Printf("quote: %s", string(q))
}

func TokenToJson(t *pay.Token) {
	j, err := json.MarshalIndent(t, "", " ")
	if err != nil {
		log.Fatalf("error marshalling response into JSON: %s \n. pay response: %v", err, t)
	}
	fmt.Printf("token: %s", string(j))
}

func TransactionToJson(r *pay.TransactionResponse) {
	j, err := json.MarshalIndent(r, "", " ")
	if err != nil {
		log.Fatalf("error marshalling response into JSON: %s \n. pay response: %v", err, r)
	}
	fmt.Printf("transaction: %s", string(j))
}

func ConfigToJson(s []byte) {

	config := &pay.ConfigData{}
	json.Unmarshal(s, &config)
	j, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		log.Fatalf("error marshalling response into JSON: %s \n. pay response: %v", err, j)
	}
	fmt.Printf("config: %s", string(j))

}

var SessionDescription = `Onramp Session command returns a unique string that can be passed into the Pay SDK to initialize the onramp widget. 
This token is associated with the Destination Wallets and App ID header provided in the request.

example:
	- pay session --address bc1qrdvlkt8rqsyj229thqzhm0q39edwdj2k7yps6x --chains "bitcoin"`

var ConfigDescription = `View a list of countries supported by Coinbase Pay and the payment methods available in each country. 

examples:

Obtain the list of countries supported by Coinbase Pay, and the payment methods available in each country.
	- pay buy --config 

Obtain the available payment options for buying Crypto with CBPay:
	- pay buy --options --country US
	- pay buy --options --country US --sub NY`

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
