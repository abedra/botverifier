package botverifier

import (
        "strings"
)

func IsUserAgentInUserAgents(useragents []string, useragent string) bool {
        if useragent == "" {
                return false
        }

        for _, test := range useragents {
                if strings.Contains(useragent, test) {
                        return true
                }
        }

        return false
}
