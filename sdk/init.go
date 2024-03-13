package sdk

import (
	"log"
	"net/http"

	"github.com/coinbase-samples/pay-sdk-go"
)

var Client *pay.Client

func InitClient() {
	creds, err := pay.SetCredentials()
	if err != nil {
		log.Fatalf("error reading environmental variable: %s", err)
	}

	Client = pay.NewClient(creds, http.Client{})
}
