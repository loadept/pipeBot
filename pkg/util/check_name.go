package util

import (
	"regexp"
)

func CheckChName(channelName string, pattern string) bool {
	matched, err := regexp.MatchString(pattern, channelName)
	if err != nil {
		return false
	}

	return matched
}
