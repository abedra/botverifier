package botverifier

import (
        "strings"
)

var GoogleUserAgents = []string{
        "Googlebot",
        "(compatible; Mediapartners-Google/2.1; +http://www.google.com/bot.html)",
        "Mediapartners-Google",
        "AdsBot-Google (+http://www.google.com/adsbot.html)",
}

type Google struct {
}

func (g *Google) IdentifiesAsBot(useragent string) bool {
        if useragent == "" {
                return false
        }

        for _, test := range GoogleUserAgents {
                if strings.Contains(useragent, test) {
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
