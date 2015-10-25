package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

type Google struct {
	UserAgentFile string
	UserAgentList []string
}

func (g *Google) LoadUserAgents() {
        file, err := os.Open(g.UserAgentFile)
        if err != nil {
                fmt.Printf("Couldn't open %s:", g.UserAgentFile, err)
                os.Exit(1)
        }
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
                g.UserAgentList = append(g.UserAgentList, scanner.Text())
        }
}

func (g *Google) IdentifiesAsBot(useragent string) bool {
        for _, test := range g.UserAgentList {
                if strings.Contains(test, useragent) {
                        return true
                }
        }

        return false
}

func (g *Google) IsBot(lookupResult []string) bool {
        for _, result := range lookupResult {
                parts := strings.Split(result, ".")
                domain := strings.Join(parts[len(parts) - 3:len(parts) - 1], ".")
                if domain == "google.com" || domain == "googlebot.com" {
                        return true
                }
        }
        return false
}

func (g *Google) Name() string {
	return "Googlebot"
}
