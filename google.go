package botverifier

import (
	"strings"
)

var UserAgents = []string{
	"Googlebot",
	"Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
	"Googlebot/2.1 (+http://www.google.com/bot.html)",
	"Googlebot-News",
	"Googlebot-Image/1.0",
	"Googlebot-Video/1.0",
	"SAMSUNG-SGH-E250/1.0 Profile/MIDP-2.0 Configuration/CLDC-1.1 UP.Browser/6.2.3.3.c.1.101 (GUI) MMP/2.0 (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
	"DoCoMo/2.0 N905i(c100;TB;W24H16) (compatible; Googlebot-Mobile/2.1; +http://www.google.com/bot.html)",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 8_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12F70 Safari/600.1.4 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
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

        for _, test := range UserAgents {
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
