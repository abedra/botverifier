package main

import (
        "fmt"
        "net"
	"strings"
	"os"
	"bufio"
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
	domain := strings.Join(strings.Split(lookupResult[0], ".")[1:3], ".")
//	fmt.Println(domain)
	if domain == "google.com" || domain == "googlebot.com" {
		return true
	} else {
		return false
	}
}

func main() {
        list := LoadGooglebotList()
	ip := "66.249.90.77"
//	ip := "50.247.150.22"
	ua := "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	if IdentifiesAsGooglebot(list, ua) {
		addr, err := net.LookupAddr(ip)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if IsGooglebot(addr) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
