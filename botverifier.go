package main

import (
        "fmt"
        "net"
        "os"
        "flag"
)

type Provider interface {
	LoadUserAgents()
        IdentifiesAsBot(useragent string) bool
        IsBot(lookupResult []string) bool
	Name() string
}

func Lookup(address string, useragent string, providers []Provider) {
        addr, err := net.LookupAddr(address)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }

        for _, provider := range providers {
                if provider.IdentifiesAsBot(useragent) {
                        fmt.Printf("%s identified as %s\n", useragent, provider.Name())
                } else {
                        fmt.Printf("%s does not identify as %s\n", useragent, provider.Name())
                }

                if provider.IsBot(addr) {
                        fmt.Printf("%s verfies as %s\n", address, provider.Name())
                } else {
                        fmt.Printf("%s cannot be verified as %s\n", address, provider.Name())
                }
        }
}

func main() {
        addressPtr   := flag.String("address", "", "IP address of requester")
        useragentPtr := flag.String("useragent", "", "User Agent of requester")
        flag.Parse()

        if *addressPtr == "" {
                fmt.Println("You must supply an IP address")
                os.Exit(1)
        }

        providers := []Provider{&Google{UserAgentFile: "googlebot.data"}}
        for _, provider := range providers {
                provider.LoadUserAgents()
        }

        if *useragentPtr != "" {
                Lookup(*addressPtr, *useragentPtr, providers)
        } else {
                Lookup(*addressPtr, "", providers)
        }
}
