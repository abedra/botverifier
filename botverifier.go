package main

import (
        "fmt"
        "net"
        "strings"
        "os"
        "bufio"
        "flag"
)

func LoadGooglebotList() []string {
        file, err := os.Open("googlebot.data")
        if err != nil {
                fmt.Println("Couldn't open googlebot.data:", err)
                os.Exit(1)
        }
        defer file.Close()

        var lines []string
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                lines = append(lines, scanner.Text())
        }

        return lines
}

func IdentifiesAsGooglebot(list []string, useragent string) bool {
        for _, test := range list {
                if strings.Contains(test, useragent) {
                        return true
                }
        }

        return false
}

func IsGooglebot(lookupResult []string) bool {
        for _, result := range lookupResult {
                parts := strings.Split(result, ".")
                domain := strings.Join(parts[len(parts) - 3:len(parts) - 1], ".")
                if domain == "google.com" || domain == "googlebot.com" {
                        return true
                }
        }
        return false
}

func Lookup(address string, useragent string, list []string) {
        if useragent == "" {
                addr, err := net.LookupAddr(address)
                if err != nil {
                        fmt.Println(err)
                        os.Exit(1)
                }
                if IsGooglebot(addr) {
                        fmt.Println("YES")
                } else {
                        fmt.Println("NO")
                }
        } else {
                if IdentifiesAsGooglebot(list, useragent) {
                        addr, err := net.LookupAddr(address)
                        if err != nil {
                                fmt.Println(err)
                                os.Exit(1)
                        }
                        if IsGooglebot(addr) {
                                fmt.Println("YES")
                        } else {
                                fmt.Println("NO")
                        }
                } else {
			fmt.Println("NO")
		}
        }
}

func main() {
        addressPtr   := flag.String("address", "", "IP address of requester")
        useragentPtr := flag.String("useragent", "", "User Agent of requester")
        flag.Parse()

        list := LoadGooglebotList()

        if *addressPtr == "" {
                fmt.Println("You must supply an IP address")
                os.Exit(1)
        }

        if *useragentPtr != "" {
                Lookup(*addressPtr, *useragentPtr, list)
        } else {
                Lookup(*addressPtr, "", list)
        }
}
