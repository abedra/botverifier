package botverifier

import (
        "strings"
)

var BingUserAgents = []string{
        "bingbot",
}

type Bing struct {
}

func (b *Bing) IdentifiesAsBot(useragent string) bool {
        return IsUserAgentInUserAgents(BingUserAgents, useragent)
}

func (b *Bing) IsBot(lookupResult []string) bool {
        for _, result := range lookupResult {
                parts := strings.Split(result, ".")
                domain := strings.Join(parts[len(parts) - 4:len(parts) - 1], ".")
                if domain == "search.msn.com" {
                        return true
                }
        }
        return false
}

func (b *Bing) Name() string {
        return "Bingbot"
}
