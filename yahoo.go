package botverifier

import (
        "strings"
	"fmt"
)

var YahooUserAgents = []string{
	"Slurp",
}

type Yahoo struct {
}

func (y *Yahoo) IdentifiesAsBot(useragent string) bool {
        if useragent == "" {
                return false
        }

        for _, test := range YahooUserAgents {
                if strings.Contains(useragent, test) {
                        return true
                }
        }

        return false
}

func (y *Yahoo) IsBot(lookupResult []string) bool {
        for _, result := range lookupResult {
                parts := strings.Split(result, ".")
                domain := strings.Join(parts[len(parts) - 3:len(parts) - 1], ".")
		fmt.Println(domain)
                if domain == "yahoo.com" {
                        return true
                }
        }

        return false
}

func (y *Yahoo) Name() string {
	return "Yahoo! Slurp"
}
