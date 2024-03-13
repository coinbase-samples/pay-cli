# Pay CLI

## Overview

The Pay CLI is a sample Command Line Interface (CLI) application that generates requests to and receives responses from Coinbase Pay's REST APIs. The Pay CLI is written in Go, using [Cobra](https://github.com/spf13/cobra) in conjunction with the [Pay Go SDK](https://github.com/coinbase-samples/pay-sdk-go).

## License

The Pay CLI is free and open source and released under the Apache License, Version 2.0.

The application and code are only available for demonstration purposes.

## Usage

To begin, navigate to your preferred directory for development and clone the Pay CLI repository and enter the directory using the following commands:

Clone repository:

```bash
git clone https://github.com/coinbase-samples pay-cli
```

Change directories:

```bash
cd pay-cli
```

Build the application:
Next, run

```bash
make start
```

to create build the application binary in the `/usr/local/bin/`

alternatively you may run:

```bash
go build -o pay
```

To verify that your application is installed correctly and accessible from any location, run the following command. It will include all available requests:

```bash
pay
```

## Examples

Obtain by config and buy options

```bash
pay buy --options --country US
```

```bash
pay buy --config
```
