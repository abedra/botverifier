package botverifier

import (
        "fmt"
        "net"
        "os"
)

type LookupResult struct {
	IdentifiesAsBot bool
	IsBot           bool
	BotName         string
}

type Provider interface {
        IdentifiesAsBot(useragent string) bool
        IsBot(lookupResult []string) bool
        Name() string
}

func Providers() []Provider {
        return []Provider{
                &Google{},
        }
}

func LookupByAddress(address string, providers []Provider) LookupResult {
        addr, err := net.LookupAddr(address)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

	result := LookupResult{IsBot: false}

        for _, provider := range providers {
                if provider.IsBot(addr) {
			result.BotName = provider.Name()
			result.IsBot = true
			return result
                }
        }

	return result
}

func LookupByAddressAndUserAgent(address string, useragent string, providers []Provider) LookupResult {
        addr, err := net.LookupAddr(address)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

	result := LookupResult{IsBot: false}

        for _, provider := range providers {
                if useragent != "" {
                        if provider.IdentifiesAsBot(useragent) {
				result.BotName = provider.Name()
				result.IdentifiesAsBot = true
                        }
                }

                if provider.IsBot(addr) {
			result.BotName = provider.Name()
			result.IsBot = true
			return result
                }
        }

	return result
}
