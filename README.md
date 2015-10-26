# botverifier

A library to determine if a requester is a known good bot.

## Setup

You can fetch this project using `go get`:

```
$ go get github.com/abedra/botverifier
```

or you can build the project directly from source:

```
$ go install
```

## Usage

The following is a complete example command line application that
demonstrates how to use the library.

```go
package main

import (
	"github.com/abedra/botverifier"
	"flag"
	"fmt"
	"os"
)

func main() {
        addressPtr   := flag.String("address", "", "IP address of requester")
        useragentPtr := flag.String("useragent", "", "User Agent of requester")
        flag.Parse()

        if *addressPtr == "" {
                fmt.Println("You must supply an IP address")
                os.Exit(1)
        }

        providers := botverifier.Providers()
	result := botverifier.LookupByAddressAndUserAgent(*addressPtr, *useragentPtr, providers)

	fmt.Println("Lookup Results for", *addressPtr)
	if result.IdentifiesAsBot || result.IsBot {
		fmt.Println("  Acting as:", result.BotName)
	}
	fmt.Println("  Identifies as a bot:", result.IdentifiesAsBot)
	fmt.Println("  Verified as a good bot:", result.IsBot)
}
```

Running this code gives you the following results:

```
$ go build test.go
$ ./test -address=50.247.150.22 -useragent="Googlebot"
Lookup Results for 50.247.150.22
  Acting as: Googlebot
  Identifies as a bot: true
  Is a bot: false
$ ./test -address=66.249.90.77 -useragent="Googlebot"
Lookup Results for 66.249.90.77
  Acting as: Googlebot
  Identifies as a bot: true
  Verified as a good bot: true
$ ./test -address=66.249.90.77
Lookup Results for 66.249.90.77
  Acting as: Googlebot
  Identifies as a bot: false
  Verified as a good bot: true
```

## Verification Information

This tool was built using verification techniques from the following provider instructions

* https://support.google.com/webmasters/answer/80553?hl=en (Googlebot)
