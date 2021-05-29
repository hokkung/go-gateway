package utils

import (
	"errors"
	"strings"
)

 
var AllowedUserAgents = []string {
	"ios",
	"android",
	"web",
}

func IsAllowedUserAgent(userAgent string) (bool, error) {
	userAgents := strings.Split(userAgent, "-")
	if len(userAgents) != 2 {
		return false, errors.New("wrong user agent")
	}

	for _, allow := range AllowedUserAgents {
		if allow == userAgents[0] {
			return true, nil
		}
	}

	return false, nil
}
