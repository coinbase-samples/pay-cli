# Pay CLI

## Overview

The Pay CLI is a sample Command Line Interface (CLI) application that generates requests to and receives responses from Coinbase Pay's REST APIs. The Pay CLI is written in Go, using Cobra in conjunction with the [Pay Go SDK](https://github.com/coinbase-samples/pay-sdk-go).

## License

The Pay CLI is free and open source and released under the Apache License, Version 2.0.

The application and code are only available for demonstration purposes.

## Usage

To begin, navigate to your preferred directory for development and clone the Pay CLI repository and enter the directory using the following commands:

Clone repository:

```bash
git clone https://github.com/coinbase-samples/pay-cli.git
```

Change directories:

```bash
cd pay-cli
```

To build the application, run the following command. This creates the application binary iand saves it in `/usr/local/bin/`:

```bash
make build
```

Alternatively, you may run the following command to build the application:

```bash
go build -o ./pay
```

To verify that your application is installed correctly and accessible from any location, run the following command. It will include all available requests:

```bash
pay
```

## Examples

Generate onramp URL

```bash
pay onramp --address 0xB6d00D83158feE6695C72ff9c5E915478A479256 --blockchains base --assets usdc
```

Example response:

```bash
Generated URL: https://pay.coinbase.com/buy/select-asset?appId=39c3d7f8-c205-463b-a54b-4279a5063977&destinationWallets=%5B%7B%22address%22%3A%220xB6d00D83158feE6695C72ff9c5E915478A479256%22%2C%22blockchains%22%3A%5B%22base%22%5D%2C%22assets%22%3A%5B%22usdc%22%5D%7D%5D
```

Obtain the available options for buying crypto with Pay.

```bash
 pay buyoptions --country US
```

Example response (truncated):

```json
{
   "id": "79691bcf-5c2d-54d0-bef9-ea0f128a472b",
   "name": "Audius",
   "symbol": "AUDIO",
   "networks": [
    {
     "name": "ethereum",
     "display_name": "Ethereum",
     "chain_id": "1",
     "contract_address": "0x18aAA7115705e8be94bfFEBDE57Af9BFc265B998"
    }
   ]
  },
  {
   "id": "79719406-aada-56c0-8194-65ef670e823b",
   "name": "Rocket Pool",
   "symbol": "RPL",
   "networks": [
    {
     "name": "ethereum",
     "display_name": "Ethereum",
     "chain_id": "1",
     "contract_address": "0xD33526068D116cE69F19A9ee46F0bd304F21A51f"
    }
   ]
  },
  {
   "id": "a0882482-a339-569c-a7cc-fa9b1c1122e4",
   "name": "Litentry",
   "symbol": "LIT",
   "networks": [
    {
     "name": "ethereum",
     "display_name": "Ethereum",
     "chain_id": "1",
     "contract_address": "0xb59490aB09A0f526Cc7305822aC65f2Ab12f9723"
    }
   ]
  }
```

Obtain the list of countries supported by Coinbase Pay, and the payment methods available in each country.

```bash
pay buyconfig
```

Obtain the list of payment methods available for **Finland**.

```bash
pay buyconfig --country FI
```

Example response (truncated):

```json
{
   "id": "d14e31fc-06e2-4f0f-a2f9-e6b48acf6b58",
   "name": "Access",
   "symbol": "ACS",
   "networks": [
    {
     "name": "solana",
     "display_name": "Solana",
     "chain_id": "0",
     "contract_address": "5MAYDfq5yxtudAhtfyuMBuHZjgAbaS9tbEyEQYAhDS5y"
    }
   ]
  },
  {
   "id": "29bb92f9-0020-49f6-af54-9e69f5ead9bb",
   "name": "Blur",
   "symbol": "BLUR",
   "networks": [
    {
     "name": "ethereum",
     "display_name": "Ethereum",
     "chain_id": "1",
     "contract_address": "0x5283D291DBCF85356A21bA090E6db59121208b44"
    }
   ]
  }
```
