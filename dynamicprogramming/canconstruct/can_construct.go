package canconstruct

import (
	"strings"
)

func CanConstructBrute(word string, wordbank []string) bool {
	if word == "" {
		return true
	}

	for _, value := range wordbank {

		if strings.HasPrefix(word, value) {
			resultString := strings.TrimPrefix(word, value)
			return CanConstructBrute(resultString, wordbank)
		}
	}

	return false
}
